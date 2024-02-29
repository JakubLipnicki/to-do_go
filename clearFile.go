package main

import (
	"log"
	"os"
)

func clearFile() error {
	err := os.Remove("tasklist.txt")
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Create("tasklist.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	return nil
}
