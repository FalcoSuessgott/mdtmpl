package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/FalcoSuessgott/mkdgo/pkg/parser"
	"github.com/caarlos0/env/v11"
	"github.com/spf13/cobra"
)

const (
	envVarPrefix        = "MDTMPL_"
	defaultTemplateFile = "README.md.tmpl"
	defaultOutputFile   = "README.md"
)

type Options struct {
	TemplateFile string
	OutputFile   string
	DryRun       bool
	Force        bool
	Version      bool
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
		Short:         "mdtmpl is a tool to render markdown templates",
		SilenceErrors: true,
		SilenceUsage:  true,
		RunE: func(cmd *cobra.Command, args []string) error {
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

	cmd.Flags().StringVarP(&o.TemplateFile, "template", "t", o.TemplateFile, "path to a mdtmpl template file")
	cmd.Flags().StringVarP(&o.OutputFile, "output", "o", o.OutputFile, "path to the output file")

	cmd.Flags().BoolVarP(&o.Force, "force", "f", o.Force, "overwrite output file")
	cmd.Flags().BoolVarP(&o.DryRun, "dry-run", "d", o.DryRun, "dry run, print output to stdout")
	cmd.Flags().BoolVar(&o.Version, "version", o.Version, "print version")

	return cmd
}

func Execute() error {
	if err := NewRootCmd().Execute(); err != nil {
		return fmt.Errorf("[ERROR] %w", err)
	}

	return nil
}
