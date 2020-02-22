package main

import "fmt"

func init() {
	ec2Cmd := &cmd{
		name:        "ec2",
		help:        "Manage ec2 instances",
		subCommands: make(map[string]*cmd),
	}

	ec2Cmd.addSubCmds(
		&cmd{
			name:    "launch",
			command: launch,
			help:    "Launch ec2 intance",
		},
		&cmd{
			name:    "terminate",
			command: terminate,
			help:    "Terminate ec2 intance",
		},
		&cmd{
			name:    "suspend",
			command: suspend,
			help:    "Suspend ec2 intance",
		},
		&cmd{
			name:    "update",
			command: update,
			help:    "Update ec2 intance",
		},
		&cmd{
			name:    "list",
			command: list,
			help:    "List ec2 intances",
		})

	registry.register(ec2Cmd)
}

func update() {
	fmt.Println("Ec2 server has been updated!")
}

func terminate() {
	fmt.Println("Ec2 server has been terminated!")
}

func suspend() {
	fmt.Println("Ec2 server has been suspended!")
}

func list() {
	fmt.Println("Ec2s have been listed!")
}

func launch() {
	fmt.Println("Ec2 server has been launched!")
}
