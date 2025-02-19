package cli_commands

type Command struct {
	name string
	args []string
}

func InitCommand(name string, args []string) Command {
	return Command{name, args}
}

func (c Command) GetName() string {
	return c.name
}

func (c Command) GetArgs() []string {
	return c.args
}
