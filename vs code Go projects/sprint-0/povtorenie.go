package main

import "fmt"

func main() {
	for i, letter := range "Hello, world!" {
		fmt.Println(i, string(letter))
	}
}
