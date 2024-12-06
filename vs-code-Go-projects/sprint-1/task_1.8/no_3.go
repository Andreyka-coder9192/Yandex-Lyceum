package task18

import (
	"fmt"
	"net/http"
	"strings"
)

func Authorization(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")

		if auth == "" {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		parts := strings.SplitN(auth, " ", 2)
		if len(parts) != 2 || parts[0] != "Basic" {
			http.Error(w, "Invalid Authorization header", http.StatusBadRequest)
			return
		}

		credentials := strings.SplitN(parts[1], ":", 2)
		if len(credentials) != 2 {
			http.Error(w, "Invalid Authorization header", http.StatusBadRequest)
			return
		}

		cuser := r.URL.Query().Get("user")
		cpass := r.URL.Query().Get("password")

		username, password := credentials[0], credentials[1]

		if username == cuser && password == cpass {
			next.ServeHTTP(w, r)
		} else {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		}
	}
}

func answerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "The answer is 42")
}

func main2() {
	http.HandleFunc("/answer/", Authorization(answerHandler))
	http.ListenAndServe(":8080", nil)
}
