package main

import (
	"go_shapes/pkg/controller"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", controller.HandleForm)
	http.HandleFunc("/shape", controller.GetShape)
	http.HandleFunc("/qsort", controller.GetSortedArray)

	log.Fatal(http.ListenAndServe(":80", nil))

}
