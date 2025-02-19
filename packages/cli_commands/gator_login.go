package cli_commands

import (
	"context"
	"fmt"
	"os"
)

func handlerLogin(s *State, cmd Command) error {
	// Check that there is only one argument for username
	if len(cmd.args) != 1 {
		return fmt.Errorf("%w, cannot use login", ErrInvalidArgument)
	}
	// Check if user is in the database
	name := cmd.args[0]
	user, err := s.DB.GetUser(context.Background(), name)
	if err != nil {
		fmt.Println("Cannot find user in db/fetching issue")
		os.Exit(1)
		return err
	}
	// Change the config of current_user
	err = s.Config.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("%w: cannot set user config to login", err)
	}
	fmt.Printf("Logged in as %s\n", user.Name)
	return nil
}
