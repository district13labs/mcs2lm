## MCS2LM cli tool
Cli tool for testing core components of MCS2LM.

### Development
**Adding new command**
- To create a new root command add a new file in the current directory
- Give a reasonable name to it, end it in cmd
- Use the `init` function to initialize the new command and its subcommands
Take a look at the code example below.
```go
func init() {
    newCmd := &cmd{
        name: "my-command",
        command: customFunction,
        help: "my-command help message",
        longHelp: "my-command really long help message",
    }

    // Register directly as root command.
    registry.register(newCmd)

    // Register as a subcommand to an existing command.
    subCmd := &cmd{
        name: "subcmd",
    }
    newCmd.addSubCmds(subCmd)
    
    // Stack the commands.
    subCmd.addSubCmds(&cmd{
        name: "another-subcmd",
    })
}

func customFunction() {
    // Magic...
}
```

Look for examples in the current implementation, e.g. `mccmd.go`.
