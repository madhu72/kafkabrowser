package server

import (
	"net/http"
	"log"
)


func ServeApp(port string) {
	fs := http.FileServer( http.Dir( "public" ) )
	log.Fatal(http.ListenAndServe( ":"+port, fs ))
}
