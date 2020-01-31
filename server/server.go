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

// Init : init the webserver
func Init(cfg Config) {	
	addr := "127.0.0.1:" + strconv.Itoa(int(cfg.Port))
	multiplexer := http.Server{
		Addr: addr,
		Handler: nil,
	}
	fmt.Printf("Menupad API version %v is running at %v\n", API_VERSION, addr)
	log.Fatal(multiplexer.ListenAndServe())
}
