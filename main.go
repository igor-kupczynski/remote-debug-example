package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	addr := ":8080"
	log.Printf("Starting server on %s\n", addr)
	handlerFunc := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		now := time.Now()
		msg := fmt.Sprintf("Hello, world at %s!\n", now.Format("2006-01-02 15:04:05"))
		_, _ = io.WriteString(w, msg)
	})
	log.Fatal(http.ListenAndServe(addr, handlerFunc))
}
