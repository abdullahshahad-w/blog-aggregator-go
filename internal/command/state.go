package command

import (
	"github.com/abdullahshahad-w/blog-aggregator-go/internal/config"
	"github.com/abdullahshahad-w/blog-aggregator-go/internal/database"
)

type State struct {
	Db *database.Queries
	Cfg *config.Config
}