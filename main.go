package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/MimiValsi/gator/internal/config"
	"github.com/MimiValsi/gator/internal/database"
)

type state struct {
	cfg *config.Config
	db  *database.Queries
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

	db, err := sql.Open("postgres", cfg.Db_url)
	if err != nil {
		log.Fatal("couldn't connect to database")
	}

	dbQueries := database.New(db)
	progState.db = dbQueries

	// ambigouos maneuver
	// cmds.register(cmd.Name, handlerLogin)
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerGetUsers)
	cmds.register("agg", handlerAgg)

	err = cmds.run(progState, cmd)
	if err != nil {
		log.Fatal(err)
	}

}
