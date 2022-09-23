package main

import (
	"go_shapes/pkg/controller"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/cmd", controller.HandleForm)
	http.HandleFunc("/shape", controller.GetShape)
	http.HandleFunc("/qsort", controller.GetSortedArray)

	log.Printf("Server started")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
