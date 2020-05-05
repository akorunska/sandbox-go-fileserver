package main

import (
    "fmt"
    "os"
)

func writeFile(contents string, filepath string) (error) {
    // todo check if such file already exists
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
    return "got file", nil
}
