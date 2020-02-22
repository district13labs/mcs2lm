package main

import (
	"fmt"
	"os"
)

var registry = cmdRegistry{
	commands: make(map[string]*cmd),
}

type cmdRegistry struct {
	commands map[string]*cmd
	// Set through build argument, CliName.
	cliName string
}

type cmdRetriever interface {
	getCommands() map[string]*cmd
	hasSubCommand() bool
}

// Fires up the command registration.
func (r *cmdRegistry) initialize(name string) {
	r.cliName = name
	if len(os.Args) < 2 {
		r.printRegisteredCmds()
		return
	}

	// Dispatch the arguments.
	err := r.mediate(os.Args[1:])
	if err != nil {
		fmt.Println(err)
	}

	return
}

// Mediates the command line arguments.
//
// cmdRegistry acts as a mediator. Calls the appropriate
// command based on the command line argument.
func (r *cmdRegistry) mediate(args []string) error {
	if args[0] == "help" {
		if len(args) < 2 {
			r.printRegisteredCmds()
			return nil
		}

		scmd, ok := r.commands[args[1]]
		if !ok {
			return fmt.Errorf("can not get help for %s: %v", args[1], errNoSuchCommand)
		}

		err := help(scmd, args[1:])
		if err != nil {
			return fmt.Errorf("can not mediate help: %v", err)
		}

		return nil
	}

	scmd, ok := r.commands[args[0]]
	if !ok {
		return fmt.Errorf("could not run root command '%s': %v", args[0], errNoSuchCommand)
	}
	if scmd.hasSubCommand() && len(args) == 1 {
		scmd.printRegisteredCmds()
		return nil
	}

	err := runCmd(scmd, args)
	if err != nil {
		return fmt.Errorf("can not mediate cmd execution: %v", err)
	}

	return nil
}

func (r *cmdRegistry) register(c *cmd) {
	registry.commands[c.name] = c
}

func (r *cmdRegistry) getCommands() map[string]*cmd {
	return r.commands
}

// Suffice helper interface
func (r *cmdRegistry) hasSubCommand() bool {
	return len(r.commands) > 0
}

func (r *cmdRegistry) printHelp() {
	r.printRegisteredCmds()
}

func (r *cmdRegistry) printRegisteredCmds() {
	b := helpHeaderBuilder("Usage of '%s':\n", r.cliName)
	for n, c := range r.commands {
		b.WriteString(fmt.Sprintf("\t%-10s\t\t %s\n", n, c.help))
	}

	fmt.Println(b.String())
}
