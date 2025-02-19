package cli_commands

import (
	"context"
	"fmt"
	"github.com/JonCyprus/BlogAggregator/internal/database"
	"time"
)

func handlerAddFeed(state *State, cmd Command) error {
	// Check for correct amount of args
	args := cmd.GetArgs()
	if len(args) != 2 {
		fmt.Println("Usage: addfeed <name> <url>")
		return ErrInvalidArgument
	}
	queries := state.GetDB()
	name := args[0]
	url := args[1]
	// Create necessary parameters and then query
	feedParams, err := InitCreateFeedParams(name, url, state)
	if err != nil {
		fmt.Println("error creating feed params for query", err)
		return err
	}
	feed, err := queries.CreateFeed(context.Background(), *feedParams)
	if err != nil {
		fmt.Println("error adding feed to database", err)
		return err
	}
	fmt.Println("Added feed to database")
	fmt.Println(feed)
	// Add to the feed_follows table
	err = handlerFollow(state, Command{
		name: "follow",
		args: []string{url}})
	if err != nil {
		return err
	}
	return nil
}

func InitCreateFeedParams(name string, url string, state *State) (*database.CreateFeedParams, error) {
	// Get UUID for currentUser from database
	currentUser := state.CurrentUser()
	queries := state.GetDB()
	currentUserData, err := queries.GetUser(context.Background(), currentUser)
	if err != nil {
		fmt.Println("Error getting user data", err)
		return nil, err
	}

	return &database.CreateFeedParams{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    currentUserData.ID,
	}, nil
}
