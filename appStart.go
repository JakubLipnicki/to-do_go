package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func appStart() {
	fmt.Println("	What you want to do?")
	fmt.Println("	Choose:")
	fmt.Println("	 add: for adding new task")
	fmt.Println("	 remove: for deleting a task")
	fmt.Println("	 display: for displaying your current tasks")
	fmt.Println("	 clear: clearing current list")
	display()
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("")
		fmt.Print("command$: ")
		scanner.Scan()

		if len(scanner.Text()) == 0 {
			continue
		}
		commandName := scanner.Text()
		trimmed := strings.TrimSpace(commandName)
		command, ok := getCommands()[trimmed]
		if ok {
			err := command.callback()
			if err != nil {
				log.Fatal(err)
			}
			continue
		} else {
			fmt.Println("not a command")
			continue
		}

	}
}

type commandCli struct {
	name     string
	callback func() error
}

func getCommands() map[string]commandCli {
	return map[string]commandCli{
		"add": {
			name:     "add",
			callback: add,
		},
		"remove": {
			name:     "remove",
			callback: remove,
		},
		"display": {
			name:     "display",
			callback: display,
		},
		"clear": {
			name:     "clear",
			callback: clearFile,
		},
	}
}
