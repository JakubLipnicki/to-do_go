package main

import (
	"bufio"
	"log"
	"os"
)

func overWrite() {

	temp, err := os.Open("temp.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer temp.Close()
	taskList, err := os.OpenFile("tasklist.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer taskList.Close()

	scanner := bufio.NewScanner(temp)

	for scanner.Scan() {
		_, err := taskList.WriteString(scanner.Text() + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
	er := os.Remove("temp.txt")
	if er != nil {
		log.Fatal(er)
	}
	file, err := os.Create("temp.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

}
