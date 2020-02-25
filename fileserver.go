package main

import (
    "log"
    "net/http"
)


func main() {
    http.HandleFunc("/", helloHandler)
    log.Fatal(http.ListenAndServe(port, nil))
}