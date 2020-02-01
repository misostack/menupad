package server

import (
	"fmt"
	"net/http"
	"html/template"
	"strconv"
	"log"
	"github.com/julienschmidt/httprouter"
)

const API_VERSION = 0.1

type Config struct {
	Port uint16
}

func GenerateHTML(w *http.ResponseWriter, title string) {
	check := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}
	t, err := template.New("webpage").Parse(MainLayout())
	check(err)
	data := struct {
		Title string
	}{
		Title: title,
	}
	err = t.Execute(*w, data)	
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	GenerateHTML(&w, "Index")
}

// type HelloHandler struct {
// 	Title string
// }

// func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	// w.Write([]byte(h.output))
// 	check := func(err error) {
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	}
// 	t, err := template.New("webpage").Parse(MainLayout())
// 	check(err)
// 	data := struct {
// 		Title string
// 	}{
// 		Title: h.Title,
// 	}
// 	err = t.Execute(w, data)
// }

// Init : init the webserver
func Init(cfg Config) {	
	addr := "127.0.0.1:" + strconv.Itoa(int(cfg.Port))
	router := httprouter.New()
	router.GET("/", Index)

	// index := HelloHandler{ Title: "Index" }
	// hello := HelloHandler{ Title: "hello" }
	// world := HelloHandler{ Title: "world" }
	// Attach handles to DefaultServerMux
	
	// http.Handle("/", &index)
	// http.Handle("/hello", &hello)
	// http.Handle("/world", &world)
	
	fmt.Printf("Menupad API version %v is running at %v\n", API_VERSION, addr)
	server := http.Server{
		Addr: addr,
		Handler: router,
	}
	log.Fatal(server.ListenAndServe())
}

func MainLayout() string{
	const tpl = `
	<!DOCTYPE html>
	<html>
		<head>
			<meta charset="UTF-8">
			<title>{{.Title}}</title>
		</head>
		<body>
			<h1>{{.Title}}</h1>
		</body>
	</html>`
	return tpl
}