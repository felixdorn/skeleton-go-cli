package main

import (
	"os"

	"github.com/owner/repository/core/handler"
)

func main() {
	os.Exit(cli.New().Run(os.Args[1:]))
}
