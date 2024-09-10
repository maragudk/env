# env

[![GoDoc](https://pkg.go.dev/badge/maragu.dev/env)](https://pkg.go.dev/maragu.dev/env)
[![Go](https://github.com/maragudk/env/actions/workflows/ci.yml/badge.svg)](https://github.com/maragudk/env/actions/workflows/ci.yml)

A small utility package to load different types of environment variables.

Made in 🇩🇰 by [maragu](https://www.maragu.dk), maker of [online Go courses](https://www.golang.dk/).

## Usage

```shell
go get maragu.dev/env
```

```go
package main

import (
  "fmt"
  "time"

  "maragu.dev/env"
)

func main() {
  env.MustLoad()
  host := env.GetStringOrDefault("HOST", "localhost")
  port := env.GetIntOrDefault("PORT", 8080)
  tls := env.GetBoolOrDefault("TLS_ENABLED", false)
  shutdownTimeout := env.GetDurationOrDefault("SHUTDOWN_TIMEOUT", time.Minute)
  fmt.Println(host, port, tls, shutdownTimeout)
}
```
