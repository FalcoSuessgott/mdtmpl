# mktmpl
This is an example `README.md.tmpl` showing the features of `mktmpl`.

## file & codeblocks
> You can include any files and wrap it in a code block `{{ file "examples/cfg.yml" | code "yml" }}`:
<!--- {{ file "examples/cfg.yml" | code "yml" }} --->
```yml
config:
  tool: mdtmpl
  files:
    - examples/cfg.yml
    - examples/complete.tmpl
    - examples/README.md

```

## exec
> You can run any commands `{{ exec "echo hallo" | code "bash" }}`:
<!--- {{ exec "echo hallo" | truncate | code "bash" }} --->
```bash
hallo
```

## sprig
> You can use all the functions from [sprig](https://masterminds.github.io/sprig/) and pipe them together `{{ "hello!" | upper | repeat 5 }}`:
<!--- {{ "hello!" | upper | repeat 5 }} --->
HELLO!HELLO!HELLO!HELLO!HELLO!

## template
> You can include other templates `{{ tmpl "examples/template.tmpl" }}`:
<!--- {{ tmpl "examples/template.tmpl" }} --->
repos:
  - repo: https://github.com/FalcoSuessgott/mdtmpl
    rev: v0.0.4
    hooks:
      - id: mdtmpl
        args: [-t=README.md.tmpl, -f, -o=README.md]


## template with vars
> You can include other templates `{{ tmplWithVars "examples/templateWithVars.tmpl" "version=v1.0.0" "name=kuberbernetes" }}`:
<!--- {{ tmplWithVars "examples/templateWithVars.tmpl" "version=v1.0.0" "name=kuberbernetes" }} --->
This is another template kuberbernetes-v1.0.0
