# go_env

```bash
go get github.com/yottab-io/go-env
```

This golang library allows retrieving environment variables with pre-defined default values, including type conversion. Additionally, the library has the ability to enforce required variables, meaning that if a variable is not defined, it will trigger an error and terminate the program.

Supports the following types:

* `string`
* `[]string`
* `bool`
* `int`
* `int64`
* `float64`

## Usage

Example:

```go
package main

import (
  "fmt"

  env "github.com/yottab-io/go-env"
)

func main() {
  // Load variable values from environment variables
  val := env.Get("HOST", "http://127.0.0.1")
  fmt.Printf("HOST=%s", val)

  // call function without default value mean is variable is Required,
  // if Not have value, it will trigger an panic(error)
  arr := env.Get("DB_PASS")
  fmt.Printf("DB_PASS=%+v", arr)

  arr := env.GetArray("APP_PORTS", []string{"80", "443"})
  fmt.Printf("APP_PORTS=%+v", arr)
}
```

```bash
$ export DB_PASS=123456
$ go run main.go
HOST=http://127.0.0.1
DB_PASS=123456
APP_PORTS=[80, 443]

$ # In this example, the environment variable is defined
$ export HOST="localhost"
$ export DB_PASS=123456
$ export APP_PORTS="8080,8443"
$ go run main.go
HOST=localhost
DB_PASS=123456
APP_PORTS=[8080, 8443]
```