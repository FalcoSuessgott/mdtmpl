# Usage
Per default, `mdtmpl` uses `README.md.tmpl` as the template file (change with `-t`) and attempts to write the rendered output to `README.md` (change with `-o`). If a `README.md` already exists, you will have to specify `--force` to overwrite its content. You can enable dry-runs using `-d`.

## CLI Args & Environment Vars

```bash
$> mdtmpl -h
template  Markdown files using Go templates and Markdown comments

Usage:
  mdtmpl [flags]

Flags:
  -d, --dry-run           dry run, print output to stdout (env: MDTMPL_DRY_RUN)
  -f, --force             overwrite output file (env: MDTMPL_FORCE)
  -h, --help              help for mdtmpl
  -o, --output string     path to the output file (env: MDTMPL_OUTPUT_FILE) (default "README.md")
  -t, --template string   path to a mdtmpl template file (env: MDTMPL_TEMPLATE_FILE) (default "README.md.tmpl")
      --version           print version (env: MDTMPL_VERSION)
```
