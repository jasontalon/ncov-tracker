package handlers

import (
	"fmt"
	"net/http"
)

func notFoundErr(w http.ResponseWriter) {
	w.WriteHeader(404)
	fmt.Fprint(w, "Not found")
}
func errIf(err error, w http.ResponseWriter) bool {
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err.Error())
		return true
	}
	return false
}
