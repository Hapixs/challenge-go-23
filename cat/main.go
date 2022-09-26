package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
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
		for _, c := range string(content) {
			z01.PrintRune(c)
		}
	}
}
