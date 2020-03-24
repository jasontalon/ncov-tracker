package internal

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/jasontalon/ncov-tracker/internal/hospital"
	"github.com/jasontalon/ncov-tracker/internal/individual"
	"github.com/jasontalon/ncov-tracker/internal/numbers"
	"github.com/jasontalon/ncov-tracker/internal/residence"
)


func TestJsonConvert(t *testing.T) {
	data, err := GetSummary()
	if err != nil {
		t.Error(err)
	}
	out := map[string]interface{}{}
	for _, row := range *data {
		out[row.Category] = row.Data
	}

	b, err := json.Marshal(out)
	if err != nil {
		t.Error(err)
		return
	}

	f, err := os.Create("json.json")

	if err != nil {
		t.Error(err)
		return
	}
	f.Write(b)

}
func TestGetSummary(t *testing.T) {
	data, err := GetSummary()
	if err != nil {
		t.Error(err)
	}
	for _, row := range *data {
		fmt.Printf("%+v\n", row)
	}
}

