package main

import (
    "fmt"
    "io/ioutil"
    "encoding/json"
    "log"
    "net/http"
    "path/filepath"
)

type FileContent struct {
  Contents string `json:"contents"`
}

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
        fmt.Fprintf(w, "post method")
      case http.MethodPost:
        fileContent := FileContent{}

        b, err := ioutil.ReadAll(r.Body)
        defer r.Body.Close()
        if err != nil {
          http.Error(w, err.Error(), 500)
          return
        }

        err = json.Unmarshal(b, &fileContent)
        if err != nil {
          http.Error(w, err.Error(), 500)
          return
        }

        output, err := json.Marshal(fileContent)
        if err != nil {
          http.Error(w, err.Error(), 500)
          return
        }
        w.Header().Set("content-type", "application/json")
        w.Write(output)
      default:
        fmt.Fprintf(w, "method not recognised")
    }
}