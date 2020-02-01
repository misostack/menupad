package server

import (
	"fmt"
	"net/http"
	"strconv"
	"log"
)

const API_VERSION = 0.1

type Config struct {
	Port uint16
}

type HelloHandler struct {}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}

// Init : init the webserver
func Init(cfg Config) {	
	addr := "127.0.0.1:" + strconv.Itoa(int(cfg.Port))
	multiplexer := http.Server{
		Addr: addr,
		Handler: nil,
	}
	hello := HelloHandler{}
	// Attach handles to DefaultServerMux
	http.Handle("/hello", &hello)
	fmt.Printf("Menupad API version %v is running at %v\n", API_VERSION, addr)
	log.Fatal(multiplexer.ListenAndServe())
}
