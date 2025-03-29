package main

import (
	"fmt"
	"os"

	"github.com/Bones1335/blog-journal-tool/internal/config"
)

type state struct {
	config *config.Config
}

func main() {
	cfg := config.Read()

	st := &state{
		config: &cfg,
	}

	cmds := commands{
		handlers: make(map[string]func(*state, command) error),
	}

	cmds.register("new", handlerNewFile)

	args := os.Args
	if len(args) < 2 {
		fmt.Println("not enough arguments")
		os.Exit(1)
	}

	commandName := args[1]
	commandArgs := args[2:]
	cmd := command{Name: commandName, Args: commandArgs}

	if err := cmds.run(st, cmd); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Exit(0)
}
