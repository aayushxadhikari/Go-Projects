package main

import (
	"flag"
	"fmt"
)

func runCLI(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: cli [create|get|list|update|delete] [options]")
		return
	}

	switch args[0]{
	case "create":
		createCMD := flag.NewFlagSet("create", flag.ExitOnError)
	}
}