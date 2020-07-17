package server

import (
	"net/http"
	"log"
)


func ServeApp(port string) {
	//fs := http.FileServer( http.Dir( "public" ) )
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(AssetFile()))
	log.Fatal(http.ListenAndServe( ":"+port, mux ))
}
