package cli_commands

import (
	"context"
	"fmt"
	"github.com/JonCyprus/BlogAggregator/internal/database"
	"strconv"
)

func handlerBrowse(state *State, cmd Command) error {
	args := cmd.GetArgs()
	if len(args) > 1 {
		fmt.Println("Usage: browse <limit[optional]>")
		return ErrInvalidArgument
	}
	var limit int32
	if len(args) == 1 {
		temp, _ := strconv.Atoi(args[0])
		limit = int32(temp)
	} else {
		limit = int32(2)
	}

	// Query database
	queries := state.GetDB()
	userData, err := queries.GetUser(context.Background(), state.CurrentUser())
	if err != nil {
		return err
	}
	posts, err := queries.GetPostsForUser(context.Background(), database.GetPostsForUserParams{userData.ID, limit})
	if err != nil {
		return err
	}
	for _, post := range posts {
		fmt.Println(post.Url, post.Description)
	}
	return nil
}
