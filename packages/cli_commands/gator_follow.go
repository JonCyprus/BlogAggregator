package cli_commands

import (
	"context"
	"fmt"
	"github.com/JonCyprus/BlogAggregator/internal/database"
	"time"
)

func handlerFollow(state *State, cmd Command) error {
	// Check argument count
	args := cmd.GetArgs()
	if len(args) != 1 {
		fmt.Println("Usage: follow <url> (to follow a feed)")
		return ErrInvalidArgument
	}
	url := args[0]

	// Input the feed follow record into the database
	feedFollowParams, err := InitCreateFeedFollowParams(state, url)
	if err != nil {
		return err
	}
	feedFollow, err := state.GetDB().CreateFeedFollow(context.Background(), *feedFollowParams)
	if err != nil {
		fmt.Println("Error creating feed follow")
		return err
	}
	fmt.Println("Successfully created feed follow")
	fmt.Println("User: ", feedFollow.UserName, "| Feed Name:", feedFollow.FeedName)
	return nil
}

func InitCreateFeedFollowParams(state *State, url string) (*database.CreateFeedFollowParams, error) {
	queries := state.GetDB()
	// Get current user data
	username := state.CurrentUser()
	userData, err := queries.GetUser(context.Background(), username)
	if err != nil {
		fmt.Println("error getting user data")
		return nil, err
	}

	// Get the feed data
	feedData, err := queries.GetFeedByURL(context.Background(), url)
	if err != nil {
		fmt.Println("error getting feed data")
		return nil, err
	}

	return &database.CreateFeedFollowParams{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    userData.ID,
		FeedID:    feedData.ID,
	}, nil
}
