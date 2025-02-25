package main

import (
	"fmt"
	"log"

	"github.com/MimiValsi/gator/internal/config"
)

func main() {
	s := new(state)
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %s\n", err)
	}
	fmt.Printf("Read config: %+v\n", cfg)

	s.Config = &cfg

	// cmds := commands{
	// 	cli: make(map[string]func(*state, command) error),
	// }

}
