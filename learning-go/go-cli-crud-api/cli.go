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

	switch args[0] {
	case "create":
		createCmd := flag.NewFlagSet("create", flag.ExitOnError)
		name := createCmd.String("name", "", "Item name")
		price := createCmd.Int("price", 0, "Item price")
		createCmd.Parse(args[1:])
		item := createItem(*name, *price)
		printJSON(item)

	case "list":
		printJSON(listItems())

	case "get":
		getCmd := flag.NewFlagSet("get", flag.ExitOnError)
		id := getCmd.Int("id", 0, "Item ID")
		getCmd.Parse(args[1:])
		if item, ok := getItem(*id); ok {
			printJSON(item)
		} else {
			fmt.Println("Item not found")
		}

	case "update":
		updateCmd := flag.NewFlagSet("update", flag.ExitOnError)
		id := updateCmd.Int("id", 0, "Item ID")
		name := updateCmd.String("name", "", "New name")
		price := updateCmd.Int("price", 0, "New price")
		updateCmd.Parse(args[1:])
		if item, ok := updateItem(*id, *name, *price); ok {
			printJSON(item)
		} else {
			fmt.Println("Item not found")
		}

	case "delete":
		deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
		id := deleteCmd.Int("id", 0, "Item ID")
		deleteCmd.Parse(args[1:])
		if ok := deleteItem(*id); ok {
			fmt.Println("Item deleted:", *id)
		} else {
			fmt.Println("Item not found")
		}

	default:
		fmt.Println("Unknown CLI command:", args[0])
	}
}
