package ytdlp

import "strings"

// Custom adds a custom flag to the command.
// The name should not contain spaces or be empty.
func (c *Command) Custom(name string, args ...string) *Command {
	if name == "" || len(strings.Split(name, " ")) > 1 {
		return c
	}

	if len(args) == 0 {
		args = nil
	}

	c.addFlag(&Flag{
		ID:   "",
		Flag: "--" + name,
		Args: args,
	})

	return c
}
