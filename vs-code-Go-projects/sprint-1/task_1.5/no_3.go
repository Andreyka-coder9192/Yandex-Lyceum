package task15

import (
	"encoding/json"
	"fmt"
)

type Student1 struct {
	Class string `json:"class"`
	Name  string `json:"name"`
}

func splitJSONByClass(jsonData []byte) (map[string][]byte, error) {
	var students []Student1
	err := json.Unmarshal(jsonData, &students)
	if err != nil {
		return nil, fmt.Errorf("Ошибка разбора JSON: %w", err)
	}

	classMap := make(map[string][]Student1)
	for _, student := range students {
		classMap[student.Class] = append(classMap[student.Class], student)
	}

	resultMap := make(map[string][]byte)
	for class, students := range classMap {
		jsonData, err := json.Marshal(students)
		if err != nil {
			return nil, fmt.Errorf("Ошибка кодирования JSON для класса %s: %w", class, err)
		}
		resultMap[class] = jsonData
	}

	return resultMap, nil
}
