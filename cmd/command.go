package main

import (
	"fmt"
)

type cmd struct {
	name        string
	command     func()
	help        string
	longHelp    string
	subCommands map[string]*cmd
}

func (c *cmd) addSubCmds(subCmd ...*cmd) {
	for _, sc := range subCmd {
		c.subCommands[sc.name] = sc
	}
}

// Suffice runner interface.
func (c *cmd) run() error {
	if c.command == nil {
		return errCommandHasNoExecutable
	}

	c.command()
	return nil
}

// Suffice helper interface.
func (c cmd) printRegisteredCmds() {
	b := helpHeaderBuilder("Usage of '%s':\n", c.name)

	for n, sc := range c.subCommands {
		b.WriteString(fmt.Sprintf("\t%-10s\t\t %s\n", n, sc.help))
	}

	fmt.Println(b.String())
}

func (c cmd) printHelp() {
	b := helpHeaderBuilder("Usage of subcommand '%s':\n", c.name)
	if c.longHelp != "" {
		b.WriteString(fmt.Sprintf("\n%s\n", c.longHelp))
	} else {
		b.WriteString(fmt.Sprintf("\n%s\n", c.help))
	}

	fmt.Println(b.String())
}

// Suffice cmdRetriever interface.
func (c cmd) getCommands() map[string]*cmd {
	return c.subCommands
}

func (c cmd) hasSubCommand() bool {
	return len(c.subCommands) > 0
}
