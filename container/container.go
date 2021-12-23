package container

import (
	"training-go-invoices/database"
	"training-go-invoices/server"
)

type Container struct {
	DataBase  *database.Database
	WebServer server.WebServer
}

func NewContainer() Container {
	database := database.Database{}

	return Container{
		DataBase: &database,

		WebServer: server.WebServer{},
	}
}
