# Testing system BMSTU IU9

# Setup local development

testsuite:

``` shell
python -m venv .venv
source .venv/bin/activate
pip install -r tests/requirements.txt
```

Pre-commit hooks:

``` shell
pip install pre-commit
pre-commit install
```

oapi-codegen:

``` shell
go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
```
