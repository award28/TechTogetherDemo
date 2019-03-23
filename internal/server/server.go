package server

import (
	"github.com/MercuryThePlanet/rest-tools"
	"github.com/award28/TechTogetherDemo/internal/root"
	"log"
	"net/http"
)

type Server struct {
	*http.Server
	port string
}

func NewServer(routers ...root.Router) *Server {
	port := "8088"

	for _, router := range routers {
		router.InitRoutes()
	}

	http.HandleFunc("/health-check",
		func(w http.ResponseWriter, r *http.Request) {
			tools.ServeJsonRes(w, http.StatusOK, "Hello World!")
		},
	)

	return &Server{&http.Server{Addr: ":" + port}, port}
}

func (srv *Server) Start() {
	log.Println("Listening on port " + srv.port)
	log.Fatal("http.ListenAndServe: ", srv.ListenAndServe())
}
