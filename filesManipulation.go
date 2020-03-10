package main

import (
    "fmt"
    "os"
)

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