package main

import (
	"log"
	"os"
	"github.com/abdullahshahad-w/blog-aggregator-go/internal/command"
	"github.com/abdullahshahad-w/blog-aggregator-go/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}	

	state := command.State{Cfg: &cfg}

	commandMap := make(map[string]func (*command.State, command.Command) error)

	commandMap["login"] = command.HandlerLogin

	args := os.Args

	if len(args) < 2 {
		log.Fatal("Usage: go run . <cmd>")
	}

	cmdStct := command.Command{
		Name: args[1],
		Args: args[2:],
	}

	if err := commandMap[cmdStct.Name](&state, cmdStct); err != nil {
		log.Fatal(err)
	}
}