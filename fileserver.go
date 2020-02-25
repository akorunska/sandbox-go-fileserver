package main

import (
    "log"
    "net/http"
)


func main() {
    http.HandleFunc("/", staticHandler)
    log.Fatal(http.ListenAndServe(port, nil))
}