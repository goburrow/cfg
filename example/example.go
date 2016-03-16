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
