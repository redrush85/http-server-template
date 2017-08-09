package main

import (
	"log"
	"net/http"
)

func main() {

	handler := Handler{
		DB:      "localhost",
		Service: &AppService{},
		// Logger as dependency?
		// Tracer as dependency?
	}

	http.Handle("/", OpenTracing(handler.Index()))
	http.Handle("/panic", RecoverMiddleware(handler.Panic()))

	err := http.ListenAndServe("0.0.0.0:9999", nil)

	if err != nil {
		log.Fatal(err)
	}

}
