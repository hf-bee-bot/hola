package main

import (
    "fmt"
    "net/http"
    "os"
)

func main() {
	fmt.Println("starting server on 8080")
    http.HandleFunc("/", HelloServer)
    http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %s! (built from %s) ", r.URL.Path[1:], os.Getenv("BRANCH"))
}
