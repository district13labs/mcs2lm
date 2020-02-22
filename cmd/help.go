package main

import (
	"fmt"
	"strings"
)

func init() {
	registry.register(&cmd{
		name:     "help",
		help:     "Displays help message",
		longHelp: fmt.Sprintf(`Run '%s help [command]' to get more info about a sub command.`, CliName),
	})
}

type helper interface {
	printHelp()
	printRegisteredCmds()
	cmdRetriever
}

func help(h helper, args []string) error {
	if h.hasSubCommand() {
		if len(args) == 1 {
			h.printRegisteredCmds()
			return nil
		}

		subCommands := h.getCommands()
		c, ok := subCommands[args[1]]
		if !ok {
			return fmt.Errorf("can not get help of '%s': %v", args[1], errNoSuchCommand)
		}

		return help(c, args[1:])
	}

	if len(args) == 1 {
		h.printHelp()
		return nil
	}

	return fmt.Errorf("invalid argument provided: %s", args[1])
}

// Builds a header and returns the strings.Builder.
//
// Displayed as follows:
//
// Usage of 'ec2':
// ---------------
//
func helpHeaderBuilder(prefix string, cmdName string) *strings.Builder {
	b := strings.Builder{}
	prefix = fmt.Sprintf(prefix, cmdName)
	b.WriteString(prefix)
	for i := 0; i < len(prefix); i++ {
		b.WriteString("-")
	}
	b.WriteString("\n")

	return &b
}
