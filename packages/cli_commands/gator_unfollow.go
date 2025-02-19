package cli_commands

import (
	"context"
	"fmt"
	"github.com/JonCyprus/BlogAggregator/internal/database"
)

func handlerUnfollow(state *State, cmd Command) error {
	args := cmd.GetArgs()
	if len(args) != 1 {
		fmt.Println("Usage: unfollow <url>")
		return ErrInvalidArgument
	}
	url := args[0]

	// Query and delete
	params, err := InitDeleteFollowFeedParams(state, url)
	if err != nil {
		return err
	}
	delRow, err := state.GetDB().DeleteFeedFollow(context.Background(), *params)
	if err != nil {
		return err
	}
	fmt.Printf("Deleted follow feed pair: %v\n", delRow)
	return nil
}

func InitDeleteFollowFeedParams(state *State, url string) (*database.DeleteFeedFollowParams, error) {
	queries := state.GetDB()
	userID, err := queries.GetUser(context.Background(), state.CurrentUser())
	if err != nil {
		return nil, err
	}
	feedID, err := queries.GetFeedByURL(context.Background(), url)
	if err != nil {
		return nil, err
	}

	return &database.DeleteFeedFollowParams{
		UserID: userID.ID,
		FeedID: feedID.ID,
	}, nil
}
