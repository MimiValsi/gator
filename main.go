package main

import (
	"fmt"
	"log"
	"os"

	"github.com/MimiValsi/gator/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %s\n", err)
	}

	// Bad initialization
	// s.Config = &cfg
	progState := &state{
		cfg: &cfg,
	}

	args := os.Args
	if len(args) < 2 {
		fmt.Println("usage: gator <cmd> [args...]")
		os.Exit(1)
	}

	cmd := command{
		Name: args[1],
		Args: args[2:],
	}

	cmds := commands{
		regCmds: make(map[string]func(*state, command) error),
	}

	// ambigouos maneuver
	// cmds.register(cmd.Name, handlerLogin)
	cmds.register("login", handlerLogin)

	err = cmds.run(progState, cmd)
	if err != nil {
		log.Fatal(err)
	}
}
