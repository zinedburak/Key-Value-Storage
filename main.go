package main

import (
	"Key_Value_Storage/controllers"
	"Key_Value_Storage/models"
	"Key_Value_Storage/routes"
	"net/http"
	"time"
)

var Db models.Store

func main() {
	Db := models.CreateStore("HelloWorld.json", 5*time.Hour)
	controllers.Db = Db
	routes.Setup()
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}
