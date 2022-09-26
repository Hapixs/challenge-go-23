package main

import (
	"io"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		io.Copy(os.Stdout, os.Stdin)
		return
	}
	for _, a := range args {
		content, err := ioutil.ReadFile(a)
		if err != nil {
			os.Stdout.WriteString("ERROR: " + err.Error() + "\n")
			os.Exit(1)
		}
		os.Stdout.WriteString(string(content))
	}
}
