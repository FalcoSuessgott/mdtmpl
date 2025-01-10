package cmd

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"

	"github.com/FalcoSuessgott/mdtmpl/pkg/template"
	"github.com/caarlos0/env/v11"
	"github.com/spf13/cobra"
)

const (
	envVarPrefix        = "MDTMPL_"
	defaultTemplateFile = "README.md.tmpl"
	defaultOutputFile   = "README.md"
)

const commentRegex = `<!---\s*(.*?)\s*--->`

var Version string

type Options struct {
	TemplateFile string `env:"TEMPLATE_FILE"`
	OutputFile   string `env:"OUTPUT_FILE"`
	DryRun       bool   `env:"DRY_RUN"`
	Force        bool   `env:"FORCE"`
	Init         bool   `env:"INIT"`
	Version      bool   `env:"VERSION"`
}

func defaultOpts() *Options {
	return &Options{
		OutputFile:   defaultOutputFile,
		TemplateFile: defaultTemplateFile,
	}
}

var initTemplate = `# ToC
<!--- {{ toc }} --->
`

// nolint: cyclop, funlen
func NewRootCmd() *cobra.Command {
	o := defaultOpts()

	if err := env.ParseWithOptions(o, env.Options{Prefix: envVarPrefix}); err != nil {
		log.Fatalf("cant parse env vars: %v", err)
	}

	cmd := &cobra.Command{
		Use:           "mdtmpl",
		Short:         "template  Markdown files using Go templates and Markdown comments",
		SilenceErrors: true,
		SilenceUsage:  true,
		RunE: func(cmd *cobra.Command, args []string) error {
			if o.Version {
				fmt.Println(Version)

				return nil
			}

			if o.Init {
				if _, err := os.Stat(o.TemplateFile); err == nil && !o.Force {
					return fmt.Errorf("template file %s already exists. Use --force to overwrite ", o.TemplateFile)
				}

				//nolint:gosec, mnd
				if err := os.WriteFile(o.TemplateFile, []byte(initTemplate), 0o644); err != nil {
					return fmt.Errorf("cannot write to %s: %w", o.TemplateFile, err)
				}
			}

			f, err := os.Open(o.TemplateFile)
			if err != nil {
				return fmt.Errorf("cannot open \"%s\": %w", o.TemplateFile, err)
			}

			defer func() {
				_ = f.Close()
			}()

			res, err := parse(f, template.WithTemplateFile(o.TemplateFile))
			if err != nil {
				return fmt.Errorf("cannot parse config %s: %w", o.TemplateFile, err)
			}

			if o.DryRun {
				fmt.Println(string(res))

				return nil
			}

			if _, err := os.Stat(o.OutputFile); err == nil && !o.Force {
				return fmt.Errorf("output file %s already exists. Use -f to overwrite", o.OutputFile)
			}

			//nolint:gosec, mnd
			if err := os.WriteFile(o.OutputFile, res, 0o644); err != nil {
				return fmt.Errorf("cannot write to %s: %w", o.OutputFile, err)
			}

			return nil
		},
	}

	cmd.Flags().StringVarP(&o.TemplateFile, "template", "t", o.TemplateFile, "path to a mdtmpl template file (env: MDTMPL_TEMPLATE_FILE)")
	cmd.Flags().StringVarP(&o.OutputFile, "output", "o", o.OutputFile, "path to the output file (env: MDTMPL_OUTPUT_FILE)")
	cmd.Flags().BoolVarP(&o.Force, "force", "f", o.Force, "overwrite output file (env: MDTMPL_FORCE)")
	cmd.Flags().BoolVarP(&o.DryRun, "dry-run", "d", o.DryRun, "dry run, print output to stdout (env: MDTMPL_DRY_RUN)")
	cmd.Flags().BoolVarP(&o.Init, "init", "i", o.Init, "Initialize a starting README.md.tmpl (env: MDTMPL_INIT)")

	cmd.Flags().BoolVar(&o.Version, "version", o.Version, "print version (env: MDTMPL_VERSION)")

	return cmd
}

func Execute() error {
	if err := NewRootCmd().Execute(); err != nil {
		return fmt.Errorf("[ERROR] %w", err)
	}

	return nil
}

func parse(r io.Reader, opts ...template.RendererOptions) ([]byte, error) {
	var resultFile bytes.Buffer

	re := regexp.MustCompile(commentRegex)

	ln := 1

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		if re.MatchString(scanner.Text()) {
			resultFile.Write(scanner.Bytes())
			resultFile.WriteString("\n")

			b := regexp.MustCompile(commentRegex).FindStringSubmatch(scanner.Text())[1]

			result, err := template.Render([]byte(b), nil, opts...)
			if err != nil {
				return nil, fmt.Errorf("cannot render template at line %d: %w", ln, err)
			}

			resultFile.Write(result.Bytes())
			resultFile.WriteString("\n")
		} else {
			resultFile.Write(scanner.Bytes())
			resultFile.WriteString("\n")
		}

		ln++
	}

	return resultFile.Bytes(), nil
}
