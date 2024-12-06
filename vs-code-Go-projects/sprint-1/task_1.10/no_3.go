package task110

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// HelloHandler обрабатывает запросы к /hello и возвращает JSON с именем.
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "Параметр 'name' отсутствует", http.StatusBadRequest)
		return
	}

	response := map[string]string{"name": name}
	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Ошибка при создании JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)

	// Логирование запроса
	log.Printf("%s %s", time.Now().Format(time.RFC3339), response)
}

func main() {
	http.HandleFunc("/hello", HelloHandler)

	fmt.Println("Сервер запущен на порту 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
