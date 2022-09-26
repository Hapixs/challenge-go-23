package main

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		io.Copy(os.Stdout, os.Stdin)
		return
	}
	for _, a := range args {
		content, err := ioutil.ReadFile(a)
		if err != nil {
			for _, c := range "ERROR: " + err.Error() {
				z01.PrintRune(c)
			}
			z01.PrintRune('\n')
			os.Exit(1)
			continue
		}
		for _, c := range string(content) {
			z01.PrintRune(c)
		}
	}
}
