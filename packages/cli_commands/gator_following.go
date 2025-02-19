package cli_commands

import (
	"context"
	"fmt"
)

func handlerFollowing(state *State, cmd Command) error {
	// Check for correct argument count
	args := cmd.GetArgs()
	if len(args) != 0 {
		fmt.Println("Usage: following (will show all your followed feeds)")
		return ErrInvalidArgument
	}

	// Fetch all follows from database
	queries := state.GetDB()
	currentUser := state.CurrentUser()
	followFeeds, err := queries.GetFeedFollowsForUser(context.Background(), currentUser)
	if err != nil {
		return fmt.Errorf("error getting feed follows: %w", err)
	}

	fmt.Printf("Following %d feeds:\n", len(followFeeds))
	for _, row := range followFeeds {
		fmt.Println(row.FeedName)
	}
	return nil
}
