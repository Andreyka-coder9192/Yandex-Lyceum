package task18

import (
	"fmt"
	"net/http"
	"regexp"
)

func Sanitize1(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		matched, _ := regexp.MatchString(`[^a-zA-Z+]`, name)
		if matched {
			w.Write([]byte("Hello, dirty hacker!"))
			return
		}
		next.ServeHTTP(w, r)
	})
}

func SetDefaultName1(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			w.Write([]byte("Hello, stranger!"))
			return
		}
		next.ServeHTTP(w, r)
	})
}

func HelloHandler1(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	greeting := fmt.Sprintf("Hello, %s!", name)
	w.Write([]byte(greeting))
}

func main1() {
	handler := Sanitize1(HelloHandler1)
	handler2 := SetDefaultName1(handler)
	http.ListenAndServe(":8000", handler2)
}
