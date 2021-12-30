package server

import (
	"net/http"
	"training-go-invoices/tools"
)

type WebServer struct {
	server http.Server
}

func (this *WebServer) CreateServer() {
	this.server = http.Server{
		Addr: ":5002",
	}

	err := this.server.ListenAndServe()
	if err != nil {
		tools.GetLogger().Printf("Error starting server: %s\n", err)
	}
}
