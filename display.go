package main

import (
	"bufio"
	"fmt"
	"os"
)

func display() error {
	// READING A FILE LINE BY LINE
	file, err := os.Open("tasklist.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	fmt.Println("")
	fmt.Println("		YOUR LIST OF TASKS:")
	i := 1
	for scanner.Scan() {
		fmt.Printf("	#%d: "+scanner.Text()+"\n", i)
		i++
	}
	file.Close()
	return nil
}
