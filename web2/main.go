package main

import (
  "fmt"
  "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello, Web2")
}

func main() {
  http.HandleFunc("/", handler) // ハンドラを登録してウェブページを表示させる
  http.HandleFunc("/2", handler) // ハンドラを登録してウェブページを表示させる
  http.ListenAndServe(":8080", nil)
}
