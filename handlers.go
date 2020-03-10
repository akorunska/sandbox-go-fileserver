package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "path/filepath"
)

func staticHandler(w http.ResponseWriter, r *http.Request) {
    // handles static html-s for the predefined paths
    fileToServe := staticPages[r.URL.Path]
    contents, err := ioutil.ReadFile(filepath.Join(staticDir, fileToServe))
    if err != nil {
        log.Printf("Unable to serve %s\n", r.URL.Path)
        return
    }
    fmt.Fprintf(w, string(contents))
}

func filesHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
       fmt.Fprintf(w, "get method")
    case http.MethodPost:
        fmt.Fprintf(w, "post method")
    default:
        fmt.Fprintf(w, "method not recognised")
    }
}