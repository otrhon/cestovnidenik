package controllers

import "github.com/otrhon/cestovnidenik/api/database"

type (
	BaseController struct {
		mongoDb database.MongoDb
	}
)

func NewBaseController(db database.MongoDb) *BaseController {
	return &BaseController{db}
}
