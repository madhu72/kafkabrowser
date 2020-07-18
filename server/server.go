package server

import (
	"net/http"
	"log"
)


func ServeApp(port string) {
	//If you want statically include html files use below lines
	//fs := http.FileServer( http.Dir( "public" ) )
	//log.Fatal(http.ListenAndServe( ":"+port, fs ))
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(AssetFile()))
	log.Fatal(http.ListenAndServe( ":"+port, mux ))
}
