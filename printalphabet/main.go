package main

import "fmt"

// simple function
// print alaphabet using ASCII ref to optimize
func main() {
	for i := 97; i <= 122; i++ {
		fmt.Printf("%c", i)
	}
	fmt.Print("\n")
}
