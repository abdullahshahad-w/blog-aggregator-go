package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/abdullahshahad-w/blog-aggregator-go/internal/command"
	"github.com/abdullahshahad-w/blog-aggregator-go/internal/config"
	"github.com/abdullahshahad-w/blog-aggregator-go/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("postgres", cfg.DbUrl)
	if err != nil {
		log.Fatal(err)
	}

	dbQueries := database.New(db)

	state := command.State{Db: dbQueries, Cfg: &cfg}

	commandsStruct := command.Commands{
		CmdMap: map[string]func (*command.State, command.Command) error{},
	}

	commandsStruct.Register("login", command.HandlerLogin)
	commandsStruct.Register("register", command.RegisterHangler)
	commandsStruct.Register("reset", command.ResetHandler)
	commandsStruct.Register("users", command.UsersHandler)
	commandsStruct.Register("agg", command.AggHandler)

	args := os.Args

	if len(args) < 2 {
		log.Fatal("Usage: go run . <cmd>")
	}

	cmdStct := command.Command{
		Name: args[1],
		Args: args[2:],
	}

	if err := commandsStruct.Run(&state, cmdStct); err != nil {
		log.Fatal(err)
	}
}