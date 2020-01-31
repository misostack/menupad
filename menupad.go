package main

import (
	menupad_server "github.com/misostack/menupad/server"
	"flag"
)



func main() {
	port := flag.Uint("port", 8081, "port")
	flag.Parse()
	cfg := menupad_server.Config{ Port: uint16(*port) }
	menupad_server.Init(cfg)
}