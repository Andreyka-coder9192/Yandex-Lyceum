package task110

import (
	"fmt"
	"os"
)

func WriteToLogFile(message string, fileName string) error {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("ошибка открытия файла: %w", err)
	}
	defer file.Close()

	file.Write([]byte("hello world"))
	//_, err = file.WriteString(message + "\n")
	/*if err != nil {
		return fmt.Errorf("ошибка записи в файл: %w", err)
	}*/

	return nil
}
