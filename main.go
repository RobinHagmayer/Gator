package main

import (
	"fmt"
	"os"

	"github.com/RobinHagmayer/Gator/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	programState := &state{cfg: &cfg}
	cmds := commands{registeredCommands: make(map[string]func(*state, command) error)}
	cmds.register("login", handlerLogin)

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Error: to few arguments")
		os.Exit(1)
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]
	cmd := command{
		name: cmdName,
		args: cmdArgs,
	}

	err = cmds.run(programState, cmd)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
