package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Holamundo(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Header.Get("Role"))
	fmt.Fprintln(w, "Hola mundo")
}

func main() {

	muxer := mux.NewRouter()

	muxer.HandleFunc("/", Holamundo)

	log.Fatal(http.ListenAndServe(":8808", muxer))
}
