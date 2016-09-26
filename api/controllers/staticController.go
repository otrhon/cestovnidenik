package controllers

import (
	"github.com/labstack/echo"
	content "github.com/otrhon/cestovnidenik/content/_go-gen"
)

type (
	StaticController struct {
		*BaseController
	}
)

func InitStaticController(bc *BaseController, e *echo.Echo) {
	sc := &StaticController{bc}

	for _, element := range content.AssetNames() {
		e.GET(element, sc.handle)
	}
}

func (sc StaticController) handle(c echo.Context) error {

	// removing leading slash (c.Path()[1:])
	x, _ := content.Asset(c.Path()[1:])
	return sc.javascript(c, x)
}
