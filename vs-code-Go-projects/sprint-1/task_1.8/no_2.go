package task18

import (
	"fmt"
	"net/http"
	"time"
)

var a, b int = 0, 1
var requestCount int = 0

func FibonacciHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%d", a)
	a, b = b, a+b
	requestCount++
}

func StartServer(t time.Duration) {
	http.HandleFunc("/", FibonacciHandler)
	http.HandleFunc("/metrics", MetricsHandler)
	http.ListenAndServe(":8080", nil)
}

func MetricsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "rpc_duration_milliseconds_count %d", requestCount)
}
