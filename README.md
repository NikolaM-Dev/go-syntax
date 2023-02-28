# Syntax: PRIVATE

## Run tests

```sh
# verbose & coverage
go test -v -cover
```

## How to open a html page with coverage report

```sh
 go test -coverprofile=coverage.out && go tool cover -html=coverage.out
```
