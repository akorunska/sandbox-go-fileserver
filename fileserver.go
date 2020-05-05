package main

import (
    "log"
    "net/http"
)


func main() {
    http.HandleFunc("/", responseHandler)
    log.Print("Serving at localhost", port)
    log.Fatal(http.ListenAndServe(port, nil))
}