package main

import "fmt"

func ConcatStringsAndInt(str1, str2 string, num int) (string, error) {
	return str1 + str2 + string(num), fmt.Errorf("division by zero is not allowed")
}

func main() {
	fmt.Println(ConcatStringsAndInt("dwa", "tri", 2))
}
