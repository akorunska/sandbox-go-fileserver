package main

import (
    "fmt"
    "io/ioutil"
    "os"
)

func fileExists(filepath string) (bool) {
    if _, err := os.Stat(filepath); err == nil {
        return true
    }
    return false
}

// file exists and read file assume that reading / writing is safe and file existence was checked by the handler

func writeFile(contents string, filepath string) (error) {
    f, err := os.Create(filepath)
    fmt.Print(err, "\n\n")
    if err != nil {
        return err
    }

    fmt.Print("file:", f)
    _, err = f.WriteString(contents)
    if err != nil {
        f.Close()
        return err
    }
    fmt.Printf("Wriring {%s} to %s\n", contents, filepath)
    err = f.Close()
    if err != nil {
        return err
    }
    return nil
}

func readFile(filepath string) (string, error) {
    content, err := ioutil.ReadFile(filepath)

    return string(content), err
}

