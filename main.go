package main

import (
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	Router *http.ServeMux
}

var server = Server{}

func main() {
	fmt.Println("Welcome Guys from docker")

	server.Run(":7878")
}

func (s *Server) initializeRoutes() {
	s.Router.HandleFunc("/", Index)
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome Guys.")
}

func (server *Server) Run(addr string) {
	server.Router = http.NewServeMux()
	server.initializeRoutes()

	fmt.Println("Start the development server at http://127.0.0.1" + addr)

	log.Fatal(http.ListenAndServe(addr, server.Router))
}
