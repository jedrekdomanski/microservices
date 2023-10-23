package main

import (
  "context"
  "net/http"
  "log"
  "os"
  "os/signal"
  "time"
  "github.com/jedrekdomanski/microservices/product-api/handlers"
  "github.com/gorilla/mux"
)

func main(){
  l := log.New(os.Stdout, "products-api", log.LstdFlags)

  productHandler := handlers.NewProducts(l)

  router := mux.NewRouter()
  getRouter := router.Methods(http.MethodGet).Subrouter()
  getRouter.HandleFunc("/products", productHandler.GetProducts)

	putRouter := router.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/products/{id:[0-9]+}", productHandler.UpdateProducts)
	putRouter.Use(productHandler.MiddlewareValidateProduct)

  postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/products", productHandler.AddProduct)
	postRouter.Use(productHandler.MiddlewareValidateProduct)

  s := &http.Server {
    Addr: ":9090",
    Handler: router,
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
