package wfs

import (
	"net/http"
)

type Server struct {
	Url    string
	Client *http.Client
}

func New(serverUri string) (*Server, error) {
	return &Server{
		Url:    serverUri,
		Client: http.DefaultClient,
	}, nil
}
