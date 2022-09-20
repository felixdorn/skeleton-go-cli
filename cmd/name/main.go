package main

import (
	"github.com/owner/repository/core/handler"
	"os"
)

func main() {
	os.Exit(cli.New().Run(os.Args[1:]))
}
