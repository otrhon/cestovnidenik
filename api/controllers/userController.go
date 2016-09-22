package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/julienschmidt/httprouter"
	"github.com/otrhon/cestovnidenik/api"
	"github.com/otrhon/cestovnidenik/api/models"
	"github.com/otrhon/cestovnidenik/views/generated-code"
)

type (
	UserController struct {
		BaseController
	}
)

func NewUserController(bc BaseController) *UserController {
	return &UserController{bc}
}

func (uc UserController) Flickr(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprint(w, tmpl.Flickr())
}

func (uc UserController) Insert(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprint(w, tmpl.Insert())
}

func (uc UserController) Save(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	record := models.Record{
		ID:      bson.NewObjectId(),
		Heading: r.FormValue("heading"),
		Content: r.FormValue("content"),
	}

	err := uc.mongoDb.InsertRecord(record)

	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, "/insert", http.StatusFound)
}

func (uc UserController) Test(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprint(w, tmpl.Index())
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		api.NotFound{}.ServeHTTP(w, r)
		return
	}

	u, err := uc.mongoDb.GetUser(id)

	if err != nil {
		api.NotFound{}.ServeHTTP(w, r)
		return
	}

	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", uj)
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := models.User{}

	json.NewDecoder(r.Body).Decode(&u)

	uc.mongoDb.InsertUser(u)

	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", uj)
}
