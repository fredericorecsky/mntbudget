package main

import (
	"log"
	"net/http"
	 _ "./endpoints/salary"
)

func main() {
	log.Fatal(http.ListenAndServe( ":8080", nil ))
}
