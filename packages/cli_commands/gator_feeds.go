package cli_commands

import (
	"context"
	"fmt"
)

func handlerFeeds(state *State, cmd Command) error {
	// Check for correct Argument count and empty string args
	if len(cmd.GetArgs()) != 0 {
		fmt.Println("Usage: feeds (gives feeds info <name> <URL> <user_creator>)")
		fmt.Println(cmd.GetArgs())
		return ErrInvalidArgument
	}

	// Query the database specifically for this function to print desired fields
	queries := state.GetDB()
	feeds, err := queries.GetFeedsForPrint(context.Background())
	if err != nil {
		fmt.Print("Error fetching feeds: ")
		return err
	}

	// Print feed info
	fmt.Println("List of feeds:")
	fmt.Println("<NAME> <URL> <USER_TO_ADD>")
	for _, feed := range feeds {
		name := feed.Name
		url := feed.Url
		user := feed.Username
		fmt.Println(name, url, user)
	}
	return nil
}
