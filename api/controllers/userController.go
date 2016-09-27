package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/labstack/echo"
	"github.com/otrhon/cestovnidenik/api/models"
	"github.com/otrhon/cestovnidenik/views/_go-gen"
)

type (
	UserController struct {
		*BaseController
	}
)

func NewUserController(bc *BaseController) *UserController {
	uc := &UserController{bc}
	return uc
}

func (uc UserController) Flickr(c echo.Context) error {
	return uc.html(c, tmpl.Flickr())
}

func (uc UserController) Insert(c echo.Context) error {
	return uc.html(c, tmpl.Insert())
}

func (uc UserController) Save(c echo.Context) error {

	record := models.Record{
		ID:      bson.NewObjectId(),
		Heading: c.FormValue("heading"),
		Content: c.FormValue("content"),
	}

	err := uc.mongoDb.InsertRecord(record)

	if err != nil {
		fmt.Println(err)
	}

	return c.Redirect(http.StatusFound, "/flickr")
}

func (uc UserController) Test(c echo.Context) error {
	return uc.html(c, tmpl.Index())
}

func (uc UserController) GetUser(c echo.Context) error {
	id := c.Param("id")

	if !bson.IsObjectIdHex(id) {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	u, err := uc.mongoDb.GetUser(id)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	uj, _ := json.Marshal(u)
	return uc.json(c, uj)
}

func (uc UserController) CreateUser(c echo.Context) error {
	u := models.User{}

	json.NewDecoder(c.Request().Body()).Decode(&u)

	uc.mongoDb.InsertUser(u)

	uj, _ := json.Marshal(u)

	return uc.json(c, uj)

}
