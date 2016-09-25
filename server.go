package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/otrhon/cestovnidenik/api"
	"github.com/otrhon/cestovnidenik/api/controllers"
	"github.com/otrhon/cestovnidenik/api/database"
)

func main() {

	r := httprouter.New()
	config := api.LoadConfig()
	database, session := database.NewMongoDb(config.MongoDb)
	defer session.Close()

	uc := controllers.NewUserController(database)

	r.NotFound = api.NotFound{}
	r.HandleMethodNotAllowed = false

	r.GET("/test", uc.Test)
	r.GET("/flickr", uc.Flickr)
	r.GET("/insert", uc.Insert)
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.POST("/save", uc.Save)

	r.ServeFiles("/content/*filepath", http.Dir("content"))
	http.ListenAndServe(":9090", r)
}
