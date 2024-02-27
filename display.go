package main

import (
	"bufio"
	"fmt"
	"os"
)

func display() {
	// READING A FILE LINE BY LINE
	file, err := os.Open("tasklist.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	fmt.Println("		YOUR LIST OF TASKS:")
	i := 1
	for scanner.Scan() {
		fmt.Printf("	#%d: "+scanner.Text()+"\n", i)
		i++
	}
}
