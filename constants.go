package main

const port = ":8080"
const staticDir = "static"
const fileStoragePrefix = "/files"

var staticPages = map[string]string{
    "/": "hello.html",
    "/test/test": "test.html",
}