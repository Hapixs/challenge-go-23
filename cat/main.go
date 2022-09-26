package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		args = append(args, os.Stdin.Name())
		return
	}
	for _, a := range args {
		content, err := ioutil.ReadFile(a)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Print(string(content))
	}
}
