package cli_commands

import "fmt"

// Storage of all the command types across files as well as the associated methods

type Commands struct {
	list map[string]func(*State, Command) error
}

func (cmds *Commands) register(name string, f func(*State, Command) error) {
	cmds.list[name] = f
	_, ok := cmds.list[name]
	if !ok {
		fmt.Printf("Error registering command: %s\n", name)
	}
	return
}

func (cmds *Commands) Run(s *State, cmd Command) error {
	cmdFunc, ok := cmds.list[cmd.name]
	if !ok {
		return fmt.Errorf("not valid commmand/command not registered: %s\n", cmd.name)
	}
	err := cmdFunc(s, cmd)
	if err != nil {
		return fmt.Errorf("error executing command %s: %w", cmd, err)
	}
	return nil
}

func (cmds *Commands) Initialize() {
	cmds.list = make(map[string]func(*State, Command) error)
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handleGetUsers)
	cmds.register("agg", handleAgg)
	cmds.register("addfeed", handlerAddFeed)
	cmds.register("feeds", handlerFeeds)
	cmds.register("following", handlerFollowing)
	cmds.register("follow", handlerFollow)
	cmds.register("unfollow", handlerUnfollow)
	cmds.register("browse", handlerBrowse)
}
