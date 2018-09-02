package server

/** REST Server **/

import (
	"fmt"
	"log"
	"mgo-n-goji/pkg"
	"net/http"
	"time"

	goji "goji.io"
)

//Server type definition
type Server struct {
	router *goji.Mux
	config *root.ServerConfig
}

//NewServer creates and returns a new server instance
func NewServer(cs root.ContactService, config *root.Config) *Server {
	s := Server{router: goji.NewMux(), config: config.Server}
	NewContactRouter(cs, s.router)
	return &s
}

//Start will apply server settings and starts it
func (s *Server) Start() {
	fmt.Println("Listening on port " + s.config.Port)
	server := &http.Server{
		Addr:         s.config.Port,
		Handler:      s.router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}
