package main

import (
	"github.com/HqOapp/translation-service/src/api"
	"github.com/HqOapp/translation-service/src/models"
)

func main() {
	models.ConnectDatabase()
	api.Serve()
}
