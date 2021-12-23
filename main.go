package main

import (
	"training-go-invoices/container"
)

func main() {
	cont := container.NewContainer()

	cont.DataBase.InitializeSQLite().Migrate().CreateSampleData()
	cont.WebServer.CreateServer()
}
