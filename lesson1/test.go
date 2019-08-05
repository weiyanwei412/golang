package main

import (
    "fmt"
    "net/http"
    "log"
)

func handler(w http.ResponseWriter, r *http.Request) {
   fmt.Fprintf(w, "hello,world!")
}

func main() {
    log.Println("Starting http server on 0.0.0.0:8888")
    http.HandleFunc("/",handler)
    http.ListenAndServe(":8888", nil)
}
