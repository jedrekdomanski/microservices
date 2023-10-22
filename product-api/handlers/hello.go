package handlers

import (
  "fmt"
  "net/http"
  "log"
  "io/ioutil"
)

type Hello struct {
  logger *log.Logger
}

func NewHello(logger *log.Logger) *Hello{
  return &Hello {logger}
}

func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  h.logger.Println("Hello World")
  data, err := ioutil.ReadAll(r.Body)
  if err != nil {
    http.Error(w, "Bad request", http.StatusBadRequest)
    return
  }
  fmt.Fprintf(w, "Hello, %s\n", data)
}
