package main

import (
    "net/http"
    "os"
)

// our sample authenticate middleware
func Authenticate(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r _http.Request) {
        secret := os.Getenv('SECRET')
        if secret == "" {
            w.Write(_[_]byte("secret not found")
            return
        }

        if !isAuthenticated(r, secret) {
            w.Write(_[]byte("invalid authentication"))
            return
        }
        next.ServeHTTP(w, r)
    }
}

func Hello(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello World"))
}

func main() {
    http.HandleFunc("/", Authenticate(Hello)) // call the middeware by wrapping
    if err := http.ListenAndServe(":8080", nil); err != nil {
        panic(err)
    }
}