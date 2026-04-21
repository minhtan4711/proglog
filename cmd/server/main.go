package main

import (
	"log"

	"github.com/minhtan4711/proglog/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8089")
	log.Fatal(srv.ListenAndServe())
}
