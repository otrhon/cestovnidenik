package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/otrhon/cestovnidenik/api/database"
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/css"
	"github.com/tdewolff/minify/html"
	"github.com/tdewolff/minify/js"
	"github.com/tdewolff/minify/json"
)

type (
	BaseController struct {
		mongoDb  *database.MongoDb
		minifier *minify.M
	}
)

func NewBaseController(db *database.MongoDb) *BaseController {
	m := minify.New()
	m.AddFunc("text/html", html.Minify)
	m.AddFunc("text/javascript", js.Minify)
	m.AddFunc("text/css", css.Minify)
	m.AddFunc("application/json", json.Minify)

	return &BaseController{db, m}
}

func (bc BaseController) minifyString(content string) string {
	x, err := bc.minifier.String("text/html", content)

	if err != nil {
		fmt.Println(err)
	}

	return x
}

func (bc BaseController) minifyBytes(content []byte, mediatype string) []byte {
	x, err := bc.minifier.Bytes(mediatype, content)

	if err != nil {
		fmt.Println(err)
	}
	return x
}

func (bc BaseController) html(c echo.Context, html string) error {
	if !c.Echo().Debug() {
		return c.HTML(http.StatusOK, bc.minifyString(html))
	}

	return c.HTML(http.StatusOK, html)
}

func (bc BaseController) json(c echo.Context, json []byte) error {
	if !c.Echo().Debug() {
		return c.JSON(http.StatusOK, bc.minifyBytes(json, "application/json"))
	}

	return c.JSON(http.StatusOK, json)
}

func (bc BaseController) css(c echo.Context, css []byte) error {
	if !c.Echo().Debug() {
		return c.Blob(http.StatusOK, "text/css", bc.minifyBytes(css, "text/css"))
	}

	return c.Blob(http.StatusOK, "text/css", css)
}

func (bc BaseController) javascript(c echo.Context, js []byte) error {
	if !c.Echo().Debug() {
		return c.Blob(http.StatusOK, "text/javascript", bc.minifyBytes(js, "text/javascript"))
	}

	return c.Blob(http.StatusOK, "text/javascript", js)
}
