package ytdlp

import "strings"

// Custom adds a custom flag to the command.
// The name should not contain spaces or be empty.
func (c *Command) Custom(name string, args ...string) *Command {
	if name == "" || len(strings.Split(name, " ")) > 1 {
		return c
	}

	var anyArgs []any
	for _, a := range args {
		anyArgs = append(anyArgs, a)
	}

	c.mu.Lock()
	c.extraFlags = append(c.extraFlags, &Flag{
		ID: "",
		Flag: "--" + name,
		Args: anyArgs,
	})
	c.mu.Unlock()

	return c
}
