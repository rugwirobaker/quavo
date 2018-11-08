package main

import (
	"github.com/rugwirobaker/quavo/cli/quavoctl/cmd"
)

func main() {
	cmd.Execute()
}

//TODO: Move the functions the client functions (i.e: Get, Set,...) to package client
//TODO : Generate a version accessible to the code on build
