package database

import (
	"fmt"

	"github.com/otrhon/cestovnidenik/api/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (
	MongoDb struct {
		session *mgo.Session
	}
)

func NewMongoDb(connectionString string) (db *MongoDb, session *mgo.Session) {
	session, err := mgo.Dial(connectionString)

	if err != nil {
		fmt.Println(err)
	}

	db = &MongoDb{session}
	return
}

func (this MongoDb) GetUser(id string) (models.User, error) {
	s := this.session.Copy()
	defer s.Close()

	oid := bson.ObjectIdHex(id)
	u := models.User{}

	err := s.DB("testing").C("users").FindId(oid).One(&u)

	return u, err
}

func (this MongoDb) InsertUser(user models.User) error {
	s := this.session.Copy()
	defer s.Close()

	user.ID = bson.NewObjectId()

	return s.DB("testing").C("users").Insert(user)
}

func (this MongoDb) InsertRecord(record models.Record) error {
	s := this.session.Copy()
	defer s.Close()

	return s.DB("testing").C("records").Insert(record)
}
