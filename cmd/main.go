package main

import (
	"flag"
	"fmt"
	"go-mail-local/internal/command"
)

func main() {

	var profile string

	flag.StringVar(&profile, "profile", "default", "Profile to use")

	flag.Parse()

	options := make(command.Options)
	options["profile"] = profile

	args := flag.Args()

	var cmd string

	if len(args) == 0 {
		cmd = "list"
		args = []string{}
	} else {
		cmd = args[0]
		args = args[1:]
	}

	kernel := command.NewKernel(options)

	if err := kernel.Run(cmd, args); err != nil {
		fmt.Println(err)
	}
}
