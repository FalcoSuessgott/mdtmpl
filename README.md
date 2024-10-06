# mdtmpl

Tired of copy-pasting your example configurations into your README? `mdtmpl` is a dead-simple little Go CLI tool that runs instructions defined in Markdown comments (e.g: `<!--- my markdown comment --->`). Under the hood `mdtmpl` uses Go templating, hence you have to surround your comments with `{{ comment }}`.

`mdtmpl` ships with the default Go-Template functions, [sprig](http://masterminds.github.io/sprig/) and some useful functions for building README files such as `code <highlighting>`, which will wrap the result of an instruction in a code block and the specified syntax highlighting.

## Example
Imagine the following `README.md.tmpl` that contains a Markdown comment `{{ file "config.yml" | code "yml" }}`, telling `mdtmpl` to include the content of `config.yml` and wrap it in a code block (`\`\`\``) using `YAML` syntax highlighting:

```
# Documentation

## Example Configuration
<!--- {{ file "config.yml" | code "yml" }} --->
```

when running `mdtmpl -t README.md.tmpl` it will write:

```
# Documentation

## Example Configuration
<!--- {{ file "Makefile" | code "make" }} --->
```make
cfg:
    setting: true
    install: false
```
```
