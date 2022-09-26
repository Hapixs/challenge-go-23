package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		reader := bufio.NewReader(os.Stdin)
		c, _ := reader.ReadString('\n')
		fmt.Print(c + "\n")
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
