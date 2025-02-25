package main

import (
	"fmt"
	"log"
	"os"

	"github.com/MimiValsi/gator/internal/config"
)

func main() {
	s := new(state)

	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %s\n", err)
	}

	s.Config = &cfg

	args := os.Args
	if len(args) < 2 {
		fmt.Println("usage: gator <cmd> <options>")
		os.Exit(1)
	}

	cmd := command{
		Name: args[1],
		Args: args[2:],
	}

	cmds := commands{
		RegCmds: make(map[string]func(*state, command) error),
	}

	cmds.register(cmd.Name, handlerLogin)
	err = cmds.run(s, cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
