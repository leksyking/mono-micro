package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	log.Println("Starting the orders microservice")

	ctx := cmd.Context()
	r, closeFn := createOrderMicroservice()

	defer closeFn()

	server := &http.Server{Addr: os.Getenv("SHOP_ORDER_SERVICE_BIND_ADDR"), Handler: r}

	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			panic(err)
		}
	}()
	<-ctx.Done()
	log.Println("closing order microservice")
	if err := server.Close(); err != nil {
		panic(err)
	}
}
