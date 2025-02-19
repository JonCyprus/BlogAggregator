package cli_commands

import "errors"

var (
	ErrCommandNotFound = errors.New("command not found")
	ErrInvalidCommand  = errors.New("command is invalid")
	ErrInvalidArgument = errors.New("invalid argument(s) for specified command")
)
