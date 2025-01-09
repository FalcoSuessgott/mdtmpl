# Installation
Find all `mdtmpl` releases [here](https://github.com/FalcoSuessgott/mdtmpl/releases) or simply download the latest by running:

## curl
```bash
version=$(curl https://api.github.com/repos/falcosuessgott/mdtmpl/releases/latest -s | jq .name -r)
curl -OL "https://github.com/FalcoSuessgott/mdtmpl/releases/download/${version}/mdtmpl_$(uname)_$(uname -m).tar.gz"
tar xzf mdtmpl_$(uname)_$(uname -m).tar.gz
chmod u+x mdtmpl
./mdtmpl version
```

## brew
```bash
brew install falcosuessgott/tap/mdtmpl
```

## docker
```bash
docker run falcosuessgott/mdtmpl
```
