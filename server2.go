package main

import (
  "net/http"
  "log"
)

func main () {
  http.Handle("/", http.FileServer(http.Dir("./static")))
  err := http.ListenAndServe(":9001", nil)
  if err != nil {
    log.Fatal("ListenAndServe: ", err)
  }
}
