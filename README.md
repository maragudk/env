# env

[![GoDoc](https://godoc.org/github.com/maragudk/env?status.svg)](https://godoc.org/github.com/maragudk/env)
[![Go](https://github.com/maragudk/env/actions/workflows/go.yml/badge.svg)](https://github.com/maragudk/env/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/maragudk/env/branch/main/graph/badge.svg)](https://codecov.io/gh/maragudk/env)

A small utility package to load different types of environment variables.

Made in 🇩🇰 by [maragu](https://www.maragu.dk), maker of [online Go courses](https://www.golang.dk/).

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
