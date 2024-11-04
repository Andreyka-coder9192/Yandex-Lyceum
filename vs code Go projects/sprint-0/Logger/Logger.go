package main

import "fmt"

func main() {
	errorLog := &Log{Level: Error}
	errorLog.Log("This is an error message")
}

// Тип для уровня логгирования
type LogLevel string

// Константы для уровней логгирования
const (
	Error LogLevel = "ERROR"
	Info  LogLevel = "INFO"
)

// Структура логгера
type Log struct {
	Level LogLevel
}

// Интерфейс для логгера
type Logger interface {
	Log(message string)
}

// Реализация метода Log для структуры Log
func (l Log) Log(message string) {
	switch l.Level {
	case Error:
		fmt.Println("ERROR:", message)
	case Info:
		fmt.Println("INFO:", message)
	default:
		fmt.Println("UNKNOWN LOG LEVEL:", message)
	}
}
