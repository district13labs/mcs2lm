# Minecraft Server Stack Launcher Module - MCS2LM
Core component for launching minecraft servers.

## System requirements

- UNIX/Linux OS
- docker
- id
- make

## Development

### Githooks
There are certain githooks provided to ensure code quality. These are located in `scripts/githooks`.  
In order to install githooks run `make setup-githooks` 

### Set up the development environment
The setup process will create a volume for permanent storage and also builds the official docker image used for development. The volume is used to store the CLI application once it has been built.

The CLI application is used only for development purposes. It's an easy way to manually test the public APIs of MCS2LM via the command line.

To set up the development environment, please run `make setup`.

### Build the CLI application
The build process will mount the volume to `/app-bin/` and builds the binary to the same path.

To build the CLI application, please run `make build`.

### Pass arguments to the container
The run script basically passes every argument inside the container. The path you're at in the beginning is `/app`.

Usage: `./scripts/app/run.sh foo bar`

### Run the CLI application
Use the run script in order to control the built CLI application. From the repository root: `./scripts/app/run.sh /app-bin/mcs2lm`

### Run tests

To check every automated tests, call the run script command from the repository root: `./scripts/app/run.sh go test ./...`

Run top-level tests matching "Foo", such as "TestFooBar": `./scripts/app/run.sh go test -run Foo `
