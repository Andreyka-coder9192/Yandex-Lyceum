package task18

import (
	"encoding/json"
	"net/http"
	"regexp"
)

type Answ struct {
	Greetings string `json:"greetings"`
	Name      string `json:"name"`
}

type Response struct {
	Status string      `json:"status"`
	Result interface{} `json:"result"`
}

func RPC(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				resp := Response{Status: "error", Result: map[string]interface{}{}}
				answe, err := json.Marshal(resp)
				if err != nil {
					http.Error(w, "Internal Server Error", http.StatusInternalServerError)
					return
				}
				w.Header().Set("Content-Type", "application/json")
				w.Write(answe)
				return
			}
		}()
		next.ServeHTTP(w, r)
	}
}

func Sanitize(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		matched, _ := regexp.MatchString(`[^a-zA-Z0-9+ ]`, name)
		if matched {
			panic("Invalid name")
		}
		next.ServeHTTP(w, r)
	}
}

func SetDefaultName(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			resp := Response{Status: "ok", Result: Answ{Greetings: "hello", Name: "stranger"}}
			answe, err := json.Marshal(resp)
			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(answe)
			return
		}
		next.ServeHTTP(w, r)
	}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "stranger"
	}
	resp := Response{Status: "ok", Result: Answ{Greetings: "hello", Name: name}}
	answe, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(answe)
}

func main() {
	handler := RPC(SetDefaultName(Sanitize(HelloHandler)))
	http.HandleFunc("/hello", handler)
	http.ListenAndServe(":8000", nil)
}
