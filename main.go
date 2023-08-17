package main

import (
	"geomys/shell"
	"os"
)

func main() {
	shell.Start(os.Stdin, os.Stdout)
}
