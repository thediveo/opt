# `opt`

[![PkgGoDev](https://pkg.go.dev/badge/github.com/thediveo/opt)](https://pkg.go.dev/github.com/thediveo/opt)
[![GitHub](https://img.shields.io/github/license/thediveo/opt)](https://img.shields.io/github/license/thediveo/opt)
![build and test](https://github.com/thediveo/opt/workflows/build%20and%20test/badge.svg?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/thediveo/opt)](https://goreportcard.com/report/github.com/thediveo/opt)
![Coverage](https://img.shields.io/badge/Coverage-100.0%25-brightgreen)

Our take on `If().Then().Else()` using Go generics. It's almost that literally...

```go
result := opt.If[string](answer == 42).Then("hooray").Else("boo!")
```


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
