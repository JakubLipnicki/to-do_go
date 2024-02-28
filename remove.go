package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func remove() error {
	fmt.Println("Enter task number that you want to delete:")
	var indexToRemove int
	fmt.Scan(&indexToRemove)

	file, err := os.Open("tasklist.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	temp, err := os.OpenFile("temp.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {
		log.Fatal(err)
	}

	defer temp.Close()

	scanner := bufio.NewScanner(file)

	i := 1

	for scanner.Scan() {

		task := string(scanner.Text())
		if i == indexToRemove {
			i++
			continue
		}
		_, err := temp.WriteString(task + "\n")
		if err != nil {
			log.Fatal(err)
		}
		i++
	}
	return nil
}
