# {{pkgTitle}}

{{pkgDesc}}

## Install

```bash
go get github.com/gookit/goutil/{{subPath}}
```

## Go Docs

- [Go docs for {{pkgName}}](https://pkg.go.dev/github.com/gookit/goutil/{{subPath}})

## Usage

{{pkgUsage}}

## Functions API

> **Note**: doc gen by run `go doc ./{{subPath}}`

```go
{{funcApis}}
```

## Code Check & Testing

```bash
gofmt -w -l ./{{subPath}}
golint ./{{subPath}}/...
```

**Testing**:

```shell
go test -v ./{{subPath}}/...
```

**Test limit by regexp**:

```shell
go test -v -run ^TestSetByKeys ./{{subPath}}/...
```

