package api

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

type Config struct {
	Debug   bool
	MongoDb string
}

func LoadConfig() (conf Config) {
	c := Config{}

	file, _ := os.Open("conf.json")
	defer file.Close()

	json.NewDecoder(file).Decode(&c)

	return c
}

func MyHTTPErrorHandler(err error, c echo.Context) {
	e := c.Echo()
	code := http.StatusInternalServerError
	msg := http.StatusText(code)
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		msg = he.Message
	}
	if e.Debug() {
		msg = err.Error()
	}
	if !c.Response().Committed() {
		if c.Request().Method() == echo.HEAD { // Issue #608
			c.NoContent(code)
		} else {
			c.String(code, msg)
		}
	}
	e.Logger().Error(err)
}
