package main

import (
  "fmt"
  "io/ioutil"
  "encoding/json"
  "log"
  "net/http"
  "path/filepath"
  "strings"
)

type FileContent struct {
  Contents string  `json:"contents"`
}

func parseFileContentFromResponce(r *http.Request) (FileContent, error) {
  fileContent := FileContent{}

  b, err := ioutil.ReadAll(r.Body)
  defer r.Body.Close()
  if err != nil {
    return FileContent{}, err
  }

  err = json.Unmarshal(b, &fileContent)
  if err != nil {
    return FileContent{}, err 
  }

  // output, err := json.Marshal(fileContent)
  // if err != nil {
  //   return nil, err 
  // }
  return fileContent, nil
}

func filesHandler(w http.ResponseWriter, r *http.Request, fileToWrite string) {
  switch r.Method {
    case http.MethodGet:
      fmt.Fprintf(w, "post method")
    case http.MethodPost:
      output, err := parseFileContentFromResponce(r)
      if err != nil {
        http.Error(w, err.Error(), 500)
      }
      err = writeFile(output.Contents, fileToWrite)
      if err != nil {
        http.Error(w, err.Error(), 500)
      }
      w.WriteHeader(200)
    default:
      fmt.Fprintf(w, "method not recognised")
  }
}

func responseHandler(w http.ResponseWriter, r *http.Request) {
  if strings.HasPrefix(r.URL.Path, fileStoragePrefix) {
    filesHandler(w, r, strings.TrimPrefix(r.URL.Path, "/"))
    return
  }

  // handles static html-s for the predefined paths
  fileToServe := staticPages[r.URL.Path]
  contents, err := ioutil.ReadFile(filepath.Join(staticDir, fileToServe))
  if err != nil {
      log.Printf("Unable to serve %s\n", r.URL.Path)
      return
  }
  fmt.Fprintf(w, string(contents))
}
