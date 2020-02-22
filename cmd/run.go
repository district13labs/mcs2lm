package main

import (
	"fmt"
)

type runner interface {
	run() error
	cmdRetriever
}

func runCmd(r runner, args []string) error {
	if r.hasSubCommand() {
		if len(args) == 1 {
			return nil
		}

		subCommands := r.getCommands()
		c, ok := subCommands[args[1]]
		if !ok {
			return fmt.Errorf("can not run command '%s': %v", args[1], errNoSuchCommand)
		}

		return runCmd(c, args[1:])
	}

	if len(args) == 1 {
		err := r.run()
		if err != nil {
			return fmt.Errorf("can not run command '%s': %v", args[0], err)
		}

		return nil
	}

	return fmt.Errorf("invalid argument provided: %s", args[1])
}
