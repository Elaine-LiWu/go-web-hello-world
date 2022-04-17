package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/demo/go-web-hello-world", func(w http.ResponseWriter, r *http.Request) {
        //fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
         fmt.Fprintf(w, "Go Web Hello World!")
    })

    http.ListenAndServe(":8081", nil)
}
