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
	name string
	args []string
}

type commands struct {
	cli map[string]func(*state, command) error
}

func handlerLogin(s *state, cmd command) error {
	if cmd.args == nil {
		return errors.New("login handler expects a single argument: <username>")
	}

	err := s.SetUser(cmd.args[0])
	if err != nil {
		return errors.New("couldn't write username to config file")
	}

	fmt.Printf("%s is login!\n", cmd.args[0])
	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.cli[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	// key -> cmd.name
	// value -> func(*state, command) error
	f, exists := c.cli[cmd.name]
	if !exists {
		return errors.New("unkown command")
	}

	err := f(s, cmd)
	if err != nil {
		return errors.New("couldn't handler function")
	}

	return nil
}
