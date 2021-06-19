# env

[![GoDoc](https://godoc.org/github.com/maragudk/env?status.svg)](https://godoc.org/github.com/maragudk/env)
[![Go](https://github.com/maragudk/env/actions/workflows/go.yml/badge.svg)](https://github.com/maragudk/env/actions/workflows/go.yml)

A small utility package to load different types of environment variables.

## Usage

```shell
go get -u github.com/maragudk/env
```

```go
package main

import (
  "fmt"
  "time"

  "github.com/maragudk/env"
)

func main() {
  host := env.GetStringOrDefault("HOST", "localhost")
  port := env.GetIntOrDefault("PORT", 8080)
  tls := env.GetBoolOrDefault("TLS_ENABLED", false)
  shutdownTimeout := env.GetDurationOrDefault("SHUTDOWN_TIMEOUT", time.Minute)
  fmt.Println(host, port, tls, shutdownTimeout)
}
```
