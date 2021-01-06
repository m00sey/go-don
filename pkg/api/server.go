package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

type Server struct {
	conf *viper.Viper
}

func (s Server) Launch() error {
	var mux = http.NewServeMux()

	mux.HandleFunc("/hello", greetingHandler)

	log.Println("listening on", fmt.Sprintf("localhost:%d", s.conf.GetInt("port")))

	return http.ListenAndServe(fmt.Sprintf("localhost:%d", s.conf.GetInt("port")), mux)
}

func New(cfg *viper.Viper) (*Server, error) {
	return &Server{
		conf: cfg,
	}, nil
}
