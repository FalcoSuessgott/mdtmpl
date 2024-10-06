package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/FalcoSuessgott/mdtmpl/pkg/parser"
	"github.com/caarlos0/env/v11"
	"github.com/spf13/cobra"
)

const (
	envVarPrefix        = "MDTMPL_"
	defaultTemplateFile = "README.md.tmpl"
	defaultOutputFile   = "README.md"
)

var Version string

type Options struct {
	TemplateFile string `env:"TEMPLATE_FILE"`
	OutputFile   string `env:"OUTPUT_FILE"`
	DryRun       bool   `env:"DRY_RUN"`
	Force        bool   `env:"FORCE"`
	Version      bool   `env:"VERSION"`
}

func defaultOpts() *Options {
	return &Options{
		OutputFile:   defaultOutputFile,
		TemplateFile: defaultTemplateFile,
	}
}

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

			f, err := os.Open(o.TemplateFile)
			if err != nil {
				return fmt.Errorf("cannot open \"%s\": %w", o.TemplateFile, err)
			}

			defer func() {
				_ = f.Close()
			}()

			res, err := parser.Parse(f)
			if err != nil {
				return fmt.Errorf("cannot parse config %s: %w", o.TemplateFile, err)
			}

			if o.DryRun {
				fmt.Println(string(res))

				return nil
			}

			if _, err := os.Stat(o.OutputFile); err == nil && !o.Force {
				return fmt.Errorf("output file %s already exists, use -f to overwrite", o.OutputFile)
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
	cmd.Flags().BoolVar(&o.Version, "version", o.Version, "print version (env: MDTMPL_VERSION)")

	return cmd
}

func Execute() error {
	if err := NewRootCmd().Execute(); err != nil {
		return fmt.Errorf("[ERROR] %w", err)
	}

	return nil
}
