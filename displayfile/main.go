package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]
	switch len(args) {
	case 0:
		fmt.Println("File name missing")
	case 1:
		content, err := ioutil.ReadFile(args[0])
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Print(string(content))
		}
	default:
		fmt.Println("Too many arguments")
	}
}
