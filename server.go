package main

import (
	"fmt"
	"net/http"

	mgo "gopkg.in/mgo.v2"

	"github.com/julienschmidt/httprouter"
	"github.com/otrhon/cestovnidenik/api"
	"github.com/otrhon/cestovnidenik/api/controllers"
)

func main() {
	r := httprouter.New()
	config := api.LoadConfig()
	s := getSession(config.MongoDb)
	uc := controllers.NewUserController(s)

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

func getSession(connectionString string) *mgo.Session {
	s, err := mgo.Dial(connectionString)

	if err != nil {
		fmt.Println("Unable to connect to MongoDb")
	}

	return s
}
