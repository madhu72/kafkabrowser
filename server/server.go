package server

import (
	"net/http"
	"log"
)

const (
	port = ":9099"
)

func ServeApp() {
	fs := http.FileServer( http.Dir( "public" ) )
	log.Fatal(http.ListenAndServe( port, fs ))
}
