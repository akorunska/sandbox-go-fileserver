package main

import (
    "log"
    "net/http"
)


func main() {
    http.HandleFunc("/", staticHandler)
    http.HandleFunc("/files", filesHandler)
    log.Fatal(http.ListenAndServe(port, nil))
}