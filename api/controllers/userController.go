package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/julienschmidt/httprouter"
	"github.com/otrhon/cestovnidenik/api"
	"github.com/otrhon/cestovnidenik/api/models"
	"github.com/otrhon/cestovnidenik/views/code"
)

type (
	UserController struct {
		session *mgo.Session
	}
)

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
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

	oid := bson.ObjectIdHex(id)

	u := models.User{}

	err := uc.session.DB("testing").C("users").FindId(oid).One(&u)

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

	u.ID = bson.NewObjectId()

	uc.session.DB("testing").C("users").Insert(u)

	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", uj)
}
