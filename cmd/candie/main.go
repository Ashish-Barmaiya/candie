package main

import (
	"os"

	"github.com/ashish-barmaiya/candie/internal/cli"
)

func main() {
	os.Exit(cli.Execute())
}
