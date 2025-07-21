package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/RobinHagmayer/Gator/internal/config"
	"github.com/RobinHagmayer/Gator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	db, err := sql.Open("postgres", cfg.DbUrl)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer db.Close()

	dbQueries := database.New(db)

	programState := &state{
		db:  dbQueries,
		cfg: &cfg,
	}

	cmds := commands{registeredCommands: make(map[string]func(*state, command) error)}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUser)
	cmds.register("agg", handlerAgg)

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
