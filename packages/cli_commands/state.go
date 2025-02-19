package cli_commands

import (
	"github.com/JonCyprus/BlogAggregator/internal/database"
	"github.com/JonCyprus/BlogAggregator/packages/config"
)

type State struct {
	Config *config.Config
	DB     *database.Queries
}

func (s *State) SetDB(queries *database.Queries) {
	s.DB = queries
}

func (s *State) SetConfig(config *config.Config) {
	s.Config = config
	return
}

func (s *State) GetConfig() *config.Config {
	return s.Config
}

func (s *State) GetDB() *database.Queries {
	return s.DB
}

func (s *State) CurrentUser() string {
	return s.Config.GetCurrentUser()
}
