package main

import "fmt"

func init() {
	mcCmd := &cmd{
		name:        "mc",
		help:        "Manage minecraft server on the ec2 instance",
		subCommands: make(map[string]*cmd),
	}

	mcCmd.addSubCmds(
		&cmd{
			name:    "start",
			command: start,
			help:    "Start a minecraft server",
		},
		&cmd{
			name:    "stop",
			command: stop,
			help:    "Stop a minecraft server",
		})

	registry.register(mcCmd)
}

func start() {
	fmt.Println("Start command has been fired!")
}

func stop() {
	fmt.Println("Stop command has been fired!")
}
