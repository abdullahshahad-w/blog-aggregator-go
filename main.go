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

	cmds := command.Commands{
		CmdMap: map[string]func (*command.State, command.Command) error{},
	}

	cmds.Register("login", command.HandlerLogin)
	cmds.Register("register", command.RegisterHangler)
	cmds.Register("reset", command.ResetHandler)
	cmds.Register("users", command.UsersHandler)
	cmds.Register("agg", command.AggHandler)
	cmds.Register("addfeed", command.MiddlewareLoggedIn(command.AddFeedHandler))
	cmds.Register("feeds", command.FeedsHandler)
	cmds.Register("follow", command.MiddlewareLoggedIn(command.FollowHandler))
	cmds.Register("following", command.MiddlewareLoggedIn(command.FollowingHandler))
	cmds.Register("unfollow", command.MiddlewareLoggedIn(command.UnfollowHandler))

	args := os.Args

	if len(args) < 2 {
		log.Fatal("Usage: go run . <cmd>")
	}

	cmdStct := command.Command{
		Name: args[1],
		Args: args[2:],
	}

	if err := cmds.Run(&state, cmdStct); err != nil {
		log.Fatal(err)
	}
}