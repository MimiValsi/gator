package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/MimiValsi/gator/internal/database"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	// cmd.Args[0] -> command name
	err := s.cfg.SetUser(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User switched successfully!")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	name := cmd.Args[0]
	timeCreation := time.Now().UTC()
	createUser := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: timeCreation,
		UpdatedAt: timeCreation,
		Name:      name,
	}

	ctx := context.Background()
	_, err := s.db.CreateUser(ctx, createUser)
	if err != nil {
		return fmt.Errorf("user already exist, %w", err)
	}

	err = s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("couldn't register current user, %w", err)
	}

	fmt.Println("User created successfully!")

	return nil
}
