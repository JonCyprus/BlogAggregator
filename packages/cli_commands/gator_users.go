package cli_commands

import (
	"context"
	"fmt"
)

func handleGetUsers(state *State, cmd Command) error {
	// Fetch users from the database
	users, err := state.DB.GetUsers(context.Background())
	if err != nil {
		fmt.Println("Error getting all users:", err)
		return err
	}

	// Get current user
	currentUser := state.CurrentUser()
	fmt.Println("Registered Users: Count ->", len(users))
	for _, user := range users {
		username := user.Name
		if username == currentUser {
			fmt.Println(username, "(current)")
		} else {
			fmt.Println(user.Name)
		}
	}
	return nil
}
