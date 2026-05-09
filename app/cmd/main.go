package main

import (
    "fmt"
    "log"
    "net/http"
    "time"
)

func Add(a, b int) int {
    return a + b
}

func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        next(w, r)
        log.Printf("method=%s path=%s duration=%s",
            r.Method, r.URL.Path, time.Since(start))
    }
}

func main() {
    http.HandleFunc("/", loggingMiddleware(func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Hello from CI/CD pipeline!")
    }))

    http.HandleFunc("/health", loggingMiddleware(func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "ok")
    }))

    log.Println("server started on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}