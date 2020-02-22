# Minecraft Server Stack Launcher Module - MCS2LM
Core component for launching minecraft servers.

## System requirements

- UNIX/Linux OS
- docker

## Development

### Githooks
There are certain githooks provided to ensure code quality. These are located in `scripts/githooks`.  
In order to install githooks run `make setup-githooks` 

You can control whether you want to run githooks inside the docker container or in the host machine.  
The default mechanism will run githooks using the container. To turn off that feature set `DOCKER` to false during the build.
```
make DOCKER=false setup-githooks
```

### Set up the development environment
The setup process will create a volume for permanent storage and also builds the official docker image used for development.  
The volume is used to store the CLI application once it has been built.

The CLI application is used only for development purposes. It's an easy way to manually test the public APIs of MCS2LM via the command line.

To set up the development environment, please run `make setup`.
The development environment also contains a helper script: `./scripts/app/run`.  
Run `./scripts/app/run` to get a detailed help message about the available commands.

### Build the CLI application
The build process will mount the volume to `/app-bin/` and builds the binary to the same path.

To build the CLI application, please run `make build`. Control build by setting `DOCKER` variable as in case of git hooks.

### Run the CLI application
The run script will wrap the cli app and executes it from inside the container.  
In order to do it so, run `./scripts/app/run mcli`.
``` 
mcli help            print available commands and a short help message  
mcli help [command]  print a detailed information of a command.  
```
For further information of how to add new commands or extend current functionality refer to `cmd/README.md`.

### Run tests
`./scripts/app/run test` will call `go test` and propagates the argument list.

### Open a shell inside the container
You can use the run script to enter the container by running `./scripts/app/run shell`.
