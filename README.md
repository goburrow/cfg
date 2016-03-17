# Configuration loader for Go applications

[![GoDoc](https://godoc.org/github.com/goburrow/cfg?status.svg)](https://godoc.org/github.com/goburrow/cfg)
[![Build Status](https://travis-ci.org/goburrow/cfg.svg?branch=master)](https://travis-ci.org/goburrow/cfg)

## Install
```
go get github.com/goburrow/cfg
```

## Example

* example.go
```go
package main

import (
	"fmt"

	"github.com/goburrow/cfg"
)

func main() {
	keys := make(map[string]string)

	// DefaultLoader loads configuration from ini file "~/.config/myApp/config"
	// if it exists and environment variables starting with "myApp_".
	loader := cfg.DefaultLoader("myApp")
	err := loader.Load(keys)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(keys)
}
```

* ~/.config/myApp/config
```ini
x = 0
y = 0

[point1]
x = 1
y = 2
```

* Run:
```
$ env myApp_y=1 myApp_point1.x=3 ./example
map[x:0 y:1 point1.x:3 point1.y:2]
```
