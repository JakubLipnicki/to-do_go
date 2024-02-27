package main

import (
	"bufio"
	"fmt"
	"os"
)

func appStart() {
	fmt.Println("	What you want to do?")
	fmt.Println("	Choose:")
	fmt.Println("	 add: for adding new task")
	fmt.Println("	 remove: for deleting a task")
	fmt.Println("	 display: for displaying your current tasks")
	fmt.Println("")
	display()
	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		fmt.Println(scanner.Text())
		add()
		remove()
	}
}
