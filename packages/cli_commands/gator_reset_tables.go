package cli_commands

import (
	"context"
	"fmt"
)

func handlerReset(state *State, cmd Command) error {
	err := resetUsers(state)
	err = resetFeeds(state)
	err = resetFeedFollows(state)
	return err
	// Check for appropriate argument count
	if len(cmd.GetArgs()) != 1 {
		fmt.Println("Incorrect argument count for reset command")
		return ErrInvalidArgument
	}

	// Check which table to reset (can add more here or create a dispatcher)
	tableToReset := cmd.GetArgs()[0]
	switch tableToReset {

	case "users":
		err := resetUsers(state)
		if err != nil {
			fmt.Println("Error trying to reset users: ", err.Error())
			return err
		}
		fmt.Println("Reset users successfully")
		return nil

	default:
		fmt.Println("Usage: reset <table_name>")
		return ErrInvalidArgument
	}
}

func resetUsers(state *State) error {
	err := state.DB.ResetUserTable(context.Background())
	if err != nil {
		fmt.Println("Error in deleting all users from table")
		return fmt.Errorf("error in deleting all users from table: %w", err)
	}
	return nil
}

func resetFeeds(state *State) error {
	err := state.DB.ResetFeedsTable(context.Background())
	if err != nil {
		fmt.Println("Error in deleting all feeds from table")
		return fmt.Errorf("error in deleting all users from table: %w", err)
	}
	return nil
}

func resetFeedFollows(state *State) error {
	err := state.DB.ResetFeedFollowsTable(context.Background())
	if err != nil {
		fmt.Println("Error in deleting all users from table")
		return fmt.Errorf("error in deleting all users from table: %w", err)
	}
	return nil
}
