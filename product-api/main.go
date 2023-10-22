package main

import (
  "context"
  "net/http"
  "log"
  "os"
  "os/signal"
  "time"
  "github.com/jedrekdomanski/microservices/product-api/handlers"
)

func main(){
  l := log.New(os.Stdout, "products-api", log.LstdFlags)

  hello_handler := handlers.NewHello(l)
  goodbye_handler := handlers.NewGoodbye(l)

  sm := http.NewServeMux()
  sm.Handle("/", hello_handler)
  sm.Handle("/goodbye", goodbye_handler)

  s := &http.Server {
    Addr: ":9090",
    Handler: sm,
    IdleTimeout: 120*time.Second,
    ReadTimeout: 1*time.Second,
    WriteTimeout: 1*time.Second,
  }

  go func(){
    l.Println("Starting server on port 9090")

    err := s.ListenAndServe()
    if err != nil {
      l.Fatal(err)
    }
  }()

  sigChan := make(chan os.Signal)
  signal.Notify(sigChan, os.Interrupt)
  signal.Notify(sigChan, os.Kill)

  sig := <- sigChan
  l.Println("Received terminate, graceful shutdown", sig)

  tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
  s.Shutdown(tc)
}
