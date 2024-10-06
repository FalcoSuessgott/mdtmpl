# mktmpl
This is an example `README.md.tmpl` showing the features of `mktmpl`.

## file
> You can include any files `{{ file "examples/cfg.yml" | code "yml" }}`:
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
> You can run any commands `{{ exec "echo hallo" }}`:
<!--- {{ exec "echo hallo" }} --->
hallo


## sprig
> You can use all the functions from [sprig](https://masterminds.github.io/sprig/) and pipe them together `{{ "hello!" | upper | repeat 5 }}`:
<!--- {{ "hello!" | upper | repeat 5 }} --->
HELLO!HELLO!HELLO!HELLO!HELLO!

## template
> You can include other templates `{{ tmpl "examples/template.tmpl" }}`:
<!--- {{ tmpl "examples/template.tmpl" }} --->
repos:
  - repo: https://github.com/FalcoSuessgott/mdtmpl
    rev: v0.0.2
    hooks:
      - id: mdtmpl
        args: [-t=README.md.tmpl, -f, -o=README.md]
