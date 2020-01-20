package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"log"
)

func main() {
	fmt.Println("MenuPad REST API serve at http://localhost:8000")
	http.HandleFunc("/", homeController)
	http.ListenAndServe(":8000", nil)
}

func homeController(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	// header should be set first at all	
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	list := []string{ "a", "b", "c", "d" }
	listJSON, err := json.Marshal(list)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		w.Write(nil)
		log.Fatal(err)
	}
	// else
	w.WriteHeader(http.StatusOK)
	w.Write(listJSON)
}