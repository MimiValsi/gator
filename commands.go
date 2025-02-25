package main

import (
	"errors"
)

type command struct {
	Name string
	Args []string
}

type commands struct {
	regCmds map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.regCmds[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	// key -> cmd.name
	// value -> func(*state, command) error
	f, ok := c.regCmds[cmd.Name]
	if !ok {
		return errors.New("command not found")
	}

	// no need to call f(...) and check for nil
	// just return the method command
	return f(s, cmd)
}
