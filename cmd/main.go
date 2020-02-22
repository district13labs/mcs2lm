package main

// CliName holds the name of the cli tool.
// This is a build time variable.
var CliName string

func main() {
	registry.initialize(CliName)
}
