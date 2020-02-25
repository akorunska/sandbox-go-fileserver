package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "path/filepath"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    // handles static html-s for the predefined paths
    fileToServe := staticPages[r.URL.Path]
    contents, err := ioutil.ReadFile(filepath.Join(staticDir, fileToServe))
    if err != nil {
        log.Printf("Unable to serve %s\n", r.URL.Path)
        return
    }
    fmt.Fprintf(w, string(contents))
}

func main() {
    http.HandleFunc("/", helloHandler)
    log.Fatal(http.ListenAndServe(port, nil))
}