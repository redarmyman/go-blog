package main

import (
    "net/http"
    "log"
    "route"
)

func main() {
    http.HandleFunc("/tags", route.TagHandler)
    http.HandleFunc("/categories", route.CategoryHandler)
    http.HandleFunc("/posts", route.PostHandler)

    log.Fatal(http.ListenAndServe(":80", nil))
}