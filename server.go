package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
	"github.com/otrhon/cestovnidenik/api"
	"github.com/otrhon/cestovnidenik/api/controllers"
	"github.com/otrhon/cestovnidenik/api/database"
)

func main() {

	config := api.LoadConfig()
	e := echo.New()
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))
	e.Use(middleware.Secure())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.SetDebug(config.Debug)
	e.SetHTTPErrorHandler(api.MyHTTPErrorHandler)

	database, session := database.NewMongoDb(config.MongoDb)
	defer session.Close()

	bc := controllers.NewBaseController(database)
	uc := controllers.NewUserController(bc)
	controllers.InitStaticController(bc, e)

	g := e.Group("/admin")
	g.Use(middleware.BasicAuth(func(username, password string) bool {
		if username == "joe" && password == "secret" {
			return true
		}
		return false
	}))

	g.GET("/test", uc.Test)
	e.GET("/flickr", uc.Flickr)
	e.GET("/insert", uc.Insert)
	e.GET("/user/:id", uc.GetUser)
	e.POST("/user", uc.CreateUser)
	e.POST("/save", uc.Save)

	e.Run(standard.New(":9090"))
}
