package task15

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name  string `json:"name"`
	Grade int    `json:"grade"`
}

func modifyJSON(jsonData []byte) ([]byte, error) {
	var students []Student

	// Распарсить JSON данные в срез структур Student
	err := json.Unmarshal(jsonData, &students)
	if err != nil {
		return nil, fmt.Errorf("Ошибка разбора JSON: %w", err)
	}

	// Добавить 1 год к полю Grade
	for i := range students {
		students[i].Grade++
	}

	// Преобразовать обновлённый срез обратно в JSON
	updatedJSON, err := json.Marshal(students)
	if err != nil {
		return nil, fmt.Errorf("Ошибка кодирования JSON: %w", err)
	}

	return updatedJSON, nil
}
