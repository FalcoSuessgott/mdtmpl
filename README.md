# mdtmpl

Tired of copy-pasting your example configurations into your README? `mdtmpl` is a dead-simple little Go CLI tool that runs instructions defined in Markdown comments (e.g: `<!--- {{ my markdown comment }} --->`).

`mdtmpl` ships with the [default Go-Template functions](https://pkg.go.dev/text/template#hdr-Functions),[sprig](http://masterminds.github.io/sprig/) and some useful functions for building README files such as:

* `{{ code <highlighting> }}`: which will wrap the result of an instruction in a code block and the specified syntax highlighting.
* `{{ exec "echo hello world" }}`: executes the given command and returns its output

You can also pipe the output of one instruction to the next one as its last parameter:

`<!--- {{ "hello!" | upper | repeat 5 }} --->` translates to `HELLO!HELLO!HELLO!HELLO!HELLO!`

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
