package main

import (
	"io"
	"os"
)

func main() {
	args := os.Args
	f, _ := os.Open(string(args[1]))
	io.Copy(os.Stdout, f)
}
