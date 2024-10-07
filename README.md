# mdtmpl
<!--- {{ file "docs/intro.tmpl" }} --->
Tired of copy-pasting your example configurations or bumping hardcoded versions in your README?

`mdtmpl` is a dead-simple little Go CLI tool that runs instructions defined in Markdown comments (e.g: `<!--- {{ my markdown comment }} --->`).

`mdtmpl` ships with the [default Go-Template functions](https://pkg.go.dev/text/template#hdr-Functions),[sprig](http://masterminds.github.io/sprig/) and some useful functions for building README files such as:

* `{{ code "<highlighting>" "<content>" }}`: which will wrap the given content in a code block using the specified syntax highlighting
* `{{ exec "<command>" }}`: executes the given command and returns its output
* `{{ tmpl "<template-file>" }}`: to include and render another template file
* `{{ tmplWithVars "<template-file>" "<key>=<value>" "<key>=<value>"}}`: to include and render another template file with additional vars in the form of  `<key>=<value>`
* `{{ conventionalCommitBump }}`: will result to the next semantic version according to the latest git commit message.

You can also pipe the output of one instruction to the next one as its last parameter:

`<!--- {{ "hello!" | upper | repeat 5 }} --->` translates to `HELLO!HELLO!HELLO!HELLO!HELLO!`

You can find a comprehensive example [here](https://github.com/FalcoSuessgott/mdtmpl/tree/main/examples) and its [`README.md.tmpl`](https://github.com/FalcoSuessgott/mdtmpl/tree/main/examples/README.md.tmpl)

Even this [`README.md`](https://github.com/FalcoSuessgott/mdtmpl/blob/main/README.md.tmpl) is generated using `mdtmpl`!

## Example
Imagine the following `README.md.tmpl` that contains a Markdown comment `{{ file "config.yml" | code "yml" }}`, telling `mdtmpl` to include the content of `config.yml` and wrap it in a code block (`` ``` ``) using `YAML` syntax highlighting:

```
### Example Configuration
Here are all available configuration options:
<!--- {{ file "config.yml" | code "yml" }} --->

### Start
You can start the installation by running `make start`
<!--- {{ exec "make start" | code "bash" }} --->
```

when running `mdtmpl -t README.md.tmpl` it will write the following output to the `README.md`:

### Example Configuration
Here are all available configuration options:
<!--- {{ file "Makefile" | code "make" }} --->
```yml
cfg:
    setting: true
    install: false
```

### Start
You can start the installation by running `make start`
<!--- {{ exec "make start" | code "bash" }} --->
```bash
[...]
```


# Installation
Find all releases [here](https://github.com/FalcoSuessgott/mdtmpl/releases) or simply download the latest by running:

```bash
version=$(curl https://api.github.com/repos/falcosuessgott/mdtmpl/releases/latest -s | jq .name -r)
curl -OL "https://github.com/FalcoSuessgott/mdtmpl/releases/download/${version}/mdtmpl_$(uname)_$(uname -m).tar.gz"
tar xzf mdtmpl_$(uname)_$(uname -m).tar.gz
chmod u+x mdtmpl
./mdtmpl version
```

# Usage
<!--- {{ exec "echo \"$> mdtmpl -h\" && go run main.go -h" | truncate | code "bash" }} --->
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

# pre-commit hook
Add the following config to your `.pre-commit-config.yaml` file and adjust the `args` to your needs.
Mae sure to run `pre-commit install` and `pre-commit autoupdate` to stay on the last version:
<!--- {{ tmpl "docs/pre-commit.tmpl" | truncate | code "yaml" }} --->
```yaml
repos:
  - repo: https://github.com/FalcoSuessgott/mdtmpl
    rev: v0.1.0
    hooks:
      - id: mdtmpl
        args: [-t=README.md.tmpl, -f, -o=README.md]
```
