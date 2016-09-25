package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type NotFound struct{}
type Config struct {
	MongoDb string
}

func (n NotFound) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	fmt.Fprint(w, "custom 404")
}

func LoadConfig() (conf Config) {
	c := Config{}

	file, _ := os.Open("conf.json")
	defer file.Close()

	json.NewDecoder(file).Decode(&c)

	return c
}
