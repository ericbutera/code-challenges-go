# code-challenges-go

code challenges, golang edition

## container registry

- [dockerhub](https://hub.docker.com/r/ericbutera/code-challenges-go) `docker pull ericbutera/code-challenges-go`
- [ghcr](https://github.com/ericbutera/code-challenges-go/pkgs/container/code-challenges-go) `docker pull ghcr.io/ericbutera/code-challenges-go:release`

## resources

- [table driven tests](https://github.com/golang/go/wiki/TableDrivenTests)
- [effective go](https://go.dev/doc/effective_go#introduction)
- [learn go with tests](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world)
- [go by example](https://gobyexample.com/)
- [testify](https://github.com/stretchr/testify)

## commands

```sh
go test
```

## notes

- [SliceTricks](https://github.com/golang/go/wiki/SliceTricks)
- In Go, := is for declaration + assignment, whereas = is for assignment only.
- no overloads
- no default parameters
  - look into <https://petomalina.medium.com/dealing-with-optional-parameters-in-go-9780f9bfbd1d>
- [no ternary](https://go.dev/doc/faq#Does_Go_have_a_ternary_form)
