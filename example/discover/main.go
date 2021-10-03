package main

import (
	"fmt"
	"path/filepath"
)
import "github.com/wulie/gokit"

func main() {
	discover, err := gokit.Discover("*", "1")
	fmt.Println(discover, err)
	matches, err := filepath.Glob("/usr/*/bin")
	fmt.Println(matches, err)
}
