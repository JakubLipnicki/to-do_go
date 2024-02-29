package main

import (
	"bufio"
	"fmt"
	"os"
)

func add() error {
	file, err := os.OpenFile("tasklist.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("What task would you like to add?")
		scanner.Scan()
		if len(scanner.Text()) == 0 {
			fmt.Println("Try again!")
			continue
		} else {
			task := string(scanner.Text())
			_, err := file.WriteString(task + "\n")
			// _, err := file.Write([]byte(scanner.Text()))
			if err != nil {
				panic(err)
			}
			fmt.Println("TASK ADDED SUCCESFULLY!")
			break
		}
	}

	return nil
}
