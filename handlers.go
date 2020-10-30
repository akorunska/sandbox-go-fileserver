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

func parseFileContentFromResponse(r *http.Request) (FileContent, error) {
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

  return fileContent, nil
}

func patchFile(w http.ResponseWriter, r *http.Request, fileToWrite string) {
  output, err := parseFileContentFromResponse(r)
  if err != nil {
    http.Error(w, err.Error(), 500)
    return
  }
  if !fileExists(fileToWrite) {
    http.Error(w, "file does not exist", 400)
    return
  }
  err = writeFile(output.Contents, fileToWrite)
  if err != nil {
    http.Error(w, err.Error(), 500)
    return
  }
  w.WriteHeader(200)
}

func postFile(w http.ResponseWriter, r *http.Request, fileToWrite string) {
  output, err := parseFileContentFromResponse(r)
  if err != nil {
    http.Error(w, err.Error(), 500)
    return
  }
  if fileExists(fileToWrite) {
    http.Error(w, "file already exists", 400)
    return
  }
  err = writeFile(output.Contents, fileToWrite)
  if err != nil {
    http.Error(w, err.Error(), 500)
    return
  }
  w.WriteHeader(200)
}

func getFile(w http.ResponseWriter, r *http.Request, filename string) {
  if !fileExists(filename) {
    http.Error(w, "file does not exist", 400)
    return
  }
  contents, err := readFile(filename)
  if err != nil {
    http.Error(w, err.Error(), 500)
    return
  }
  w.WriteHeader(200)
  fmt.Fprintf(w, contents)
}

func filesHandler(w http.ResponseWriter, r *http.Request, filename string) {
  switch r.Method {
    case http.MethodGet:
      getFile(w, r, filename)
    case http.MethodPost:
      postFile(w, r, filename)
    case http.MethodPatch:
      patchFile(w, r, filename)
    default:
      fmt.Fprintf(w, "method not recognised")
  }
}

func responseHandler(w http.ResponseWriter, r *http.Request) {
  // todo check basic authorization

  if strings.HasPrefix(r.URL.Path, fileStoragePrefix) {
    filesHandler(w, r, strings.TrimPrefix(r.URL.Path, "/"))
    return
  }

  // handles static html-s for the predefined paths
  fileToServe := staticPages[r.URL.Path]
  contents, err := ioutil.ReadFile(filepath.Join(staticDir, fileToServe))
  if err != nil {
      log.Printf("unable to serve %s\n", r.URL.Path)
      return
  }
  fmt.Fprintf(w, string(contents))
}
