package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jasontalon/ncov-tracker/internal/hospital"
	"github.com/jasontalon/ncov-tracker/internal/individual"
	"github.com/jasontalon/ncov-tracker/internal/numbers"
	"github.com/jasontalon/ncov-tracker/internal/residence"
)

func Hospital(w http.ResponseWriter, r *http.Request) {
	data, err := hospital.GetData()
	if errIf(err, w) {
		return
	}
	render(data, w)
}

func Individual(w http.ResponseWriter, r *http.Request) {
	data, err := individual.GetData()
	if errIf(err, w) {
		return
	}
	render(data, w)
}

func Numbers(w http.ResponseWriter, r *http.Request) {
	data, err := numbers.GetData()
	if errIf(err, w) {
		return
	}
	render(data, w)
}
func Residence(w http.ResponseWriter, r *http.Request) {
	data, err := residence.GetData()
	if errIf(err, w) {
		return
	}
	render(data, w)
}

func render(data interface{}, w http.ResponseWriter) {
	b, err := json.MarshalIndent(data, "", "\t")
	if errIf(err, w) {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(*&b))
}
