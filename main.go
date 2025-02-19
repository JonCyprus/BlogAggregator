package main

import (
	"database/sql"
	"fmt"
	"github.com/JonCyprus/BlogAggregator/internal/database"
	"github.com/JonCyprus/BlogAggregator/packages/cli_commands"
	"github.com/JonCyprus/BlogAggregator/packages/config"
	_ "github.com/lib/pq"
	"os"
	"strings"
)

func main() {
	// Take the command line arguments
	args := os.Args
	//args = clean_args(args)

	// Check that there is at least a command name argument
	if len(args) < 2 {
		fmt.Println("Usage: BlogAggregator <command> <args>")
		os.Exit(1)
	}
	// Initialize the state of the program from config
	appConfig, err := config.ReadConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var state cli_commands.State
	state.SetConfig(&appConfig)

	// Open Database connection
	db, err := sql.Open("postgres", state.Config.ConfigGetDBURL())
	if err != nil {
		fmt.Println(err)
	}
	dbQueries := database.New(db)
	state.SetDB(dbQueries)

	// Initialize the needed Command Struct
	commandName := strings.ToLower(args[1])
	var commandArgs []string
	if len(args) > 2 {
		commandArgs = args[2:]
	} else {
		commandArgs = []string{}
	}
	command := cli_commands.InitCommand(commandName, commandArgs)

	// Initialize the valid commands
	var commands cli_commands.Commands
	commands.Initialize()

	// Run the command
	err = commands.Run(&state, command)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	os.Exit(0)
}

// Helper function to clean arguments
func clean_args(args []string) []string {
	for i, arg := range args {
		args[i] = strings.ToLower(arg)
	}
	return args
}
