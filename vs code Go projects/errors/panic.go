package main

import (
	"fmt"
)

func main() {
	fmt.Println("Начало")
	panic("Что-то пошло не так")
	fmt.Println("Конец")
}

func exampleFunction() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from panic:", r)
        }
    }()