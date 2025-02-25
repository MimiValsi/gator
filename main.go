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

	cmds := commands{
		cli: make(map[string]func(*state, command) error),
	}

	args := os.Args
	if len(args) < 2 {
		fmt.Println("usage: gator <cmd> <options>")
		os.Exit(-1)
	}

	if args[1] == "login" {
		cmds.cli[args[1]] = handlerLogin
	}
}
