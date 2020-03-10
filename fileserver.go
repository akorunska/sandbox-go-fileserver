package main

import (
    "log"
    "net/http"
)


func main() {
    http.HandleFunc("/", responseHandler)
    log.Fatal(http.ListenAndServe(port, nil))
}