package main

import (
	"flag"
	"fmt"
	"log"
	"mime"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func main() {
	startServer()
}
func startServer() {
	portPtr := flag.Int("p", 8081, "Port number to run the server on")
	flag.Parse()
	port := *portPtr
	mr := mux.NewRouter()
	mr.HandleFunc("/whoami", whoamiHandler)
	go mime.AddExtensionType(".css", "text/css; charset=utf-8")
	go mime.AddExtensionType(".js", "application/javascript; charset=utf-8")
	//Setup a static router for HTML/CSS/JS
	mr.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./resources"))))

	log.Printf("Listening for requests on :%v\n",port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), mr)
}

func whoamiHandler(w http.ResponseWriter, r *http.Request) {
	rhost := strings.Split(r.RemoteAddr,":")[0]
	w.Write([]byte(rhost))
}
