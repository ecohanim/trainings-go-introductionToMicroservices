package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"./handlers"
)

// rest server app
// run app with "go run main.go"
// test product with:
// 	get    - "curl localhost:9090 | jq"
// 	create - "curl -v localhost:9090 -d '{}' | jq"
//			 "curl -v localhost:9090 -X POST -d '{"name": "tea", "description": "a nice cup of tea"}'"
// 	update - "curl -v localhost:9090/1 -X PUT -d '{"name": "tea", "description": "a nice cup of tea"}'"
// 	delete - "curl -v localhost:9090/1 -X DELETE | jq" (not implemented)
// debug product with "curl -v localhost:9090"
func main() {
	l := log.New(os.Stdout, "Debug: ", log.LstdFlags)

	// create the handlers
	ph := handlers.NewProducts(l)

	// create a new serve mux and register the handlers
	sm := http.NewServeMux()
	sm.Handle("/", ph)

	// create a new server
	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// http.ListenAndServe(":9090", sm)
	l.Printf("Starting Server on port %s\n", s.Addr)

	// start the server
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Recieved terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
