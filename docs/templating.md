# Templating
A basic `mdtmpl` instruction looks like this:

```yaml
<!--- {{ <template-function> <arg01> <args02> ... }} --->
```

`mdtmpl` parses the template file and all its markdown comments and renders its instructions. It uses the [Go`s Template Engine](https://pkg.go.dev/text/template).

Follow this document to see which template functions are supported.

## Piping
You can [pipe the output of one instruction to the next template function as its **last argument**](https://pkg.go.dev/text/template#hdr-Pipelines):

```yaml
<!--- {{ <template-function> <arg01> <args02> | <template-function> <args> ... }} --->
```

For example: `<!--- {{ "hello!" | upper | repeat 5 }} --->` will result in:
`HELLO!HELLO!HELLO!HELLO!HELLO!`.

## Template Functions
`mdtmpl` includes all [`sprig`](http://masterminds.github.io/sprig/) and [Go`s predefined template functions](https://pkg.go.dev/text/template#hdr-Functions).

Furthermore, the following functions are also available:

### `code "<language>" "<content>"`
> Syntax highlight a given content in a [specified language](https://github.com/github-linguist/linguist/blob/main/lib/linguist/languages.yml) within a code block.

=== "`README.md.tmpl`"
    ```c
    <!--- {{ echo "this is a command" | code "bash" }} --->
    ```

=== "`README.md`"
    <!--- {{ echo "this is a command" | code "bash" }} --->
    ```bash
    this is a command
    ```

### `exec "<command>"`
> Executes the given command

!!! tip
    `truncate` removes any trailing empty lines. Useful after `exec`

=== "`README.md.tmpl`"
    ```yaml
    <!--- {{ exec "echo hello world" | truncate | code "bash" }} --->
    ```

=== "`README.md`"
    <!--- {{ exec "echo hello world" | truncate | code "bash" }} --->
    ```bash
    hello world
    ```

### `file "<path>"`
> Includes the content of the given file

```yaml
# settings.yaml
settings:
    basic_auth: false
```

=== "`README.md.tmpl`"
    ```c
    <!--- {{ file "settings.yml" | code "yaml" }} --->
    ```

=== "`README.md`"
    <!--- {{ file "settings.yml" | code "yaml" }} --->
    ```yaml
    settings:
        basic_auth: false
    ```

### `tmpl "<template-file>"`
> Includes the rendered content of the given template

```yaml
# docs/template.tmpl
This is a test {{ exec "echo template" }}
```

=== "`README.md.tmpl`"
    ```c
    <!--- {{ tmpl "docs/template.tmpl" }} --->
    ```

=== "`README.md`"
    <!--- {{ tmpl "docs/template.tmpl" }} --->
    This is a test template

### `tmplWithVars "<template-file>" "<key=<value>" "<key=<value>"`
> Includes the rendered content of the given template

```yaml
# docs/template.tmpl
This is another template {{ .name }}-{{ .version }}
```

=== "`README.md.tmpl`"
    ```c
    <!--- {{ tmplWithVars "docs/template.tmpl" "version=v1.0.0" "name=kuberbernetes" }} --->
    ```

=== "`README.md`"
    <!--- {{ tmplWithVars "testdata/template.tmpl" "version=v1.0.0" "name=kuberbernetes" }} --->
    This is another template kuberbernetes-v1.0.0

### `stripansi "<content>"`
> Strips any Color Codes from a given content

!!! tip
    Useful when a command outputs colored output

=== "`README.md.tmpl`"
    ```c
    <!--- {{ exec "docker ps" | stripansi | code "bash" }} --->
    ```

=== "`README.md`"
    <!--- {{ exec "docker ps" | stripansi | code "bash" }} --->
    ```bash
    CONTAINER ID   IMAGE        COMMAND                  CREATED       STATUS          PORTS                                       NAMES
    cf4f9cec8faa   registry:2   "/entrypoint.sh /etc…"   7 weeks ago   Up 29 minutes   0.0.0.0:5000->5000/tcp, :::5000->5000/tcp   registry
    ```

### `toc`
> Inserts a Markdown Table of Content

!!! note
    For now it does not work for any headings that are included after `toc` function invocation. For example when using `file` or `tmpl`/`tmplWithVars`

=== "`README.md.tmpl`"
    ```c
    # ToC
    <!--- {{ toc }} --->
    # 1. Heading
    ## 2. Heading
    ### 3. Heading
    ## 4. Heading
    ```

=== "`README.md`"
    ```
    # ToC
    - [ToC](#toc)
    - [1. Heading](#1.-heading)
    - [2. Heading](#2.-heading)
        - [3. Heading](#3.-heading)
    - [4. Heading](#4.-heading)

    # 1. Heading
    ## 2. Heading
    ### 3. Heading
    ## 4. Heading
    ```
