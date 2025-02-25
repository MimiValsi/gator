package main

import (
	"errors"
	"fmt"

	"github.com/MimiValsi/gator/internal/config"
)

type state struct {
	*config.Config
}

type command struct {
	Name string
	Args []string
}

type commands struct {
	regCmds map[string]func(*state, command) error
}

func handlerLogin(s *state, cmd command) error {
	if cmd.Args == nil {
		return errors.New("login handler expects a single argument: <username>")
	}

	err := s.SetUser(cmd.Args[0])
	if err != nil {
		return errors.New("couldn't write username to config file")
	}

	fmt.Printf("%s is login!\n", cmd.Args[0])
	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.regCmds[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	// key -> cmd.name
	// value -> func(*state, command) error
	f, exists := c.regCmds[cmd.Name]
	if !exists {
		return errors.New("unkown command")
	}

	if len(cmd.Args) == 0 {
		return errors.New("username is required")
	}

	err := f(s, cmd)
	if err != nil {
		return errors.New("couldn't handler function")
	}

	return nil
}
