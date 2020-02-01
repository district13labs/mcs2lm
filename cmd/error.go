package main

import (
	"errors"
)

var errNoSuchCommand = errors.New("command doesn't exist")
var errCommandHasNoExecutable = errors.New("command has no executable")
