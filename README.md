# `opt`

[![PkgGoDev](https://pkg.go.dev/badge/github.com/thediveo/opt)](https://pkg.go.dev/github.com/thediveo/opt)
[![GitHub](https://img.shields.io/github/license/thediveo/opt)](https://img.shields.io/github/license/thediveo/opt)
![build and test](https://github.com/thediveo/opt/actions/workflows/buildandtest.yaml/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/thediveo/opt)](https://goreportcard.com/report/github.com/thediveo/opt)
![Coverage](https://img.shields.io/badge/Coverage-100.0%25-brightgreen)

Our take on `If().Then().Else()` using Go generics. It's almost that literally...

### v2

Requires Go 1.27 and later with the new generics methods language feature. The
syntactic benefit of v2 over v1 is that there is no need anymore to explicitly
specify the value type to `opt.If`. Instead, the value type is derived from
`Then(v)`, automatically type-checking `Else(v)` to be of the same type.

```go
//go:build go1.27
result := opt.If(answer == 42).Then("hooray").Else("boo!")
```

### v1

```go
//go:build !go1.27
result := opt.If[string](answer == 42).Then("hooray").Else("boo!")
```

According to [godbolt.org's Compiler Explorer](https://godbolt.org/z/Wdj34rz1q)
the Go compiler gives the opt package short shrift: as optimized as writing the
traditional elaborate `var x T; if (cond) { x = ... } else { x = ... }`.

> [!NOTE] In the hopefully near future, when [spec: generic methods for
> Go](https://github.com/golang/go/issues/77273) has landed, we should be able
> to drop the type parameter from the `If` type.

## DevContainer

> [!CAUTION]
>
> Do **not** use VSCode's "~~Dev Containers: Clone Repository in Container
> Volume~~" command, as it is utterly broken by design, ignoring
> `.devcontainer/devcontainer.json`.

1. `git clone https://github.com/thediveo/opt`
2. in VSCode: Ctrl+Shift+P, "Dev Containers: Open Workspace in Container..."
3. select `opt.code-workspace` and off you go...

## Contributing

Please see [CONTRIBUTING.md](CONTRIBUTING.md).

## Copyright and License

`opt` is Copyright 2026 Harald Albrecht, and licensed under the Apache License,
Version 2.0.
