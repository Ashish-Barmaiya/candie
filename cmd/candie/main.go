package main

import (
	"fmt"
	"os"

	"github.com/ashish-barmaiya/candie/internal/cli"
	"github.com/ashish-barmaiya/candie/internal/config"
)

func main() {
	cfg := config.DefaultConfig()
	fmt.Println(cfg)

	os.Exit(cli.Execute())
}
