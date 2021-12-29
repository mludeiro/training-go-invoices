package main

import (
	"training-go-invoices/container"
)

func main() {
	cont := container.NewContainer()

	cont.DataBase.InitializeMySQL().Migrate().CreateSampleData()
	cont.WebServer.CreateServer()
}
