package main

import (
	"net/http"

	mgo "gopkg.in/mgo.v2"

	"github.com/julienschmidt/httprouter"
	"github.com/otrhon/cestovnidenik/api"
	"github.com/otrhon/cestovnidenik/api/controllers"
)

func main() {
	r := httprouter.New()
	s := getSession()
	uc := controllers.NewUserController(s)

	r.NotFound = api.NotFound{}
	r.HandleMethodNotAllowed = false

	r.GET("/test", uc.Test)
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	http.ListenAndServe("localhost:9090", r)
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost")

	if err != nil {
		panic(err)
	}

	return s
}
