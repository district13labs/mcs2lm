# Minecraft Server Stack Launcher Module - MCS2LM
Core component for launching minecraft servers.

## System requirements

- UNIX/Linux OS
- docker

## Development

### Githooks
There are certain githooks provided to ensure code quality. These are located in `scripts/githooks`.  
In order to install githooks run `make setup-githooks` 

You can control whether you want to run githooks inside the docker container or in the host machine. To use githooks inside the container, keep `DOCKER="true"` in the `.env` file as it is, otherwise change it to anything else, for example `DOCKER="false"`.

### Set up the development environment
The setup process will create a volume for permanent storage and also builds the official docker image used for development. The volume is used to store the CLI application once it has been built.

The CLI application is used only for development purposes. It's an easy way to manually test the public APIs of MCS2LM via the command line.

To set up the development environment, please run `make setup`.

### Build the CLI application
The build process will mount the volume to `/app-bin/` and builds the binary to the same path.

To build the CLI application, please run `make build`.

### Run the CLI application
Use the run script in order to control the built CLI application. From the repository root: `./scripts/app/run mcli`

For more information about what you can do with the run script, please run `./scripts/app/run` and you'll see a help message.

### Run tests
`./scripts/app/run test` will call `go test` and passes every arguments to it.

### Open a shell inside the container
You can use the run script to go inside the container running `./scripts/app/run shell`
