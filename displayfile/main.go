package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("File name missing")
		return
	} else if len(args) > 1 {
		fmt.Println("Too many arguments")
		return
	}
	content, err := ioutil.ReadFile("thermopylae.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content))
}
