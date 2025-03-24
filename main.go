package main

import (
	"fmt"
	"os"

	"github.com/Bones1335/blog-tool/internal/config"
)

type state struct {
	config *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("editor: %v\nsave_location: %v\n", cfg.Editor, cfg.SaveLocation)
}
