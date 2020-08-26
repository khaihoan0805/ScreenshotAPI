package main

import (
	"log"
	"fmt"
	"net/http"
)

func homepageHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "hello world")
	checkError(err)
}

func  checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
}
func main() {
	http.HandleFunc("/", homepageHandler)
	fmt.Println("server is listening on port 3000")
	log.Panic(
		http.ListenAndServe(":3000", nil)
	)
}
