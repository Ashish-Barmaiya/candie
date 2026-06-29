package main

import (
	"fmt"

	"github.com/ashish-barmaiya/candie/internal/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", cfg)
}
