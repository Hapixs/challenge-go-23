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
		fmt.Println(map[bool]string{true: string(content), false: err.Error()})
	default:
		fmt.Println("Too many arguments")
	}
}
