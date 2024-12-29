# pre-commit hook
Add the following config to your `.pre-commit-config.yaml` file and adjust the `args` to your needs.
Mae sure to run `pre-commit install` and `pre-commit autoupdate` to stick to the latest version:
```yaml
repos:
  - repo: https://github.com/FalcoSuessgott/mdtmpl
    rev: v0.0.6
    hooks:
      - id: mdtmpl
        args: [-t=README.md.tmpl, -f, -o=README.md]
```
