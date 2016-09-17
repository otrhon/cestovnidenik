package api

import (
	"fmt"
	"net/http"
)

type NotFound struct{}

func (n NotFound) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	fmt.Fprint(w, "custom 404")
}
