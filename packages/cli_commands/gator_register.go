package cli_commands

import (
	"context"
	"fmt"
	"github.com/JonCyprus/BlogAggregator/internal/database"
	"github.com/google/uuid"
	"os"
	"strings"
	"time"
)

func handlerRegister(state *State, cmd Command) error {
	// Check that only one argument ("name") is passed
	args := cmd.GetArgs()
	if len(args) != 1 {
		return ErrInvalidArgument
	}
	// Check that name is not empty
	name := args[0]
	if strings.TrimSpace(name) == "" {
		return fmt.Errorf("name cannot be blank")
	}

	// Create the params needed for the query and execture the query
	UserParams, err := InitCreateUserParams(name)
	if err != nil {
		return err
	}
	newUser, err := state.DB.CreateUser(context.Background(), UserParams)
	if err != nil {
		fmt.Println("duplicate user, cannot create")
		os.Exit(1)
		return err
	}
	err = state.Config.SetUser(newUser.Name)
	if err != nil {
		return err
	}
	fmt.Println("user created")
	fmt.Println(newUser)
	return nil
}

func InitCreateUserParams(name string) (database.CreateUserParams, error) {
	return database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name}, nil
}
