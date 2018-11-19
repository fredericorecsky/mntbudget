package main

import (
	"log"
	"net/http"
	 _ "mntbudget/endpoints/salary"
)

func main() {
	// if google cloud uses 8080
	log.Fatal(http.ListenAndServe( ":8081", nil ))
}
