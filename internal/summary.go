package internal

import (
	"fmt"

	"github.com/jasontalon/ncov-tracker/internal/hospital"
	"github.com/jasontalon/ncov-tracker/internal/individual"
	"github.com/jasontalon/ncov-tracker/internal/numbers"
	"github.com/jasontalon/ncov-tracker/internal/residence"
)

func GetAll() (*[]AllocSummary, error) {
	ch := make(chan AllocSummary, 4)
	m := map[string]func() (interface{}, error){
		"hospital": func() (interface{}, error) {
			return hospital.GetData()
		},
		"individual": func() (interface{}, error) {
			return individual.GetData()
		},
		"numbers": func() (interface{}, error) {
			return numbers.GetData()
		},
		"residence": func() (interface{}, error) {
			return residence.GetData()
		},
	}

	for key, f := range m {
		go func(cc chan AllocSummary, f func() (interface{}, error), category string) {
			data, err := f()
			cc <- AllocSummary{Category: category, Data: data, Error: err}
		}(ch, f, key)
	}

	dataset := []AllocSummary{}
	for i := 0; 4 > i; i++ {
		dataset = append(dataset, <-ch)
	}

	for _, data := range dataset {
		if data.Error != nil {
			return nil, fmt.Errorf("%s: %s\n", data.Category, data.Error.Error())
		}
	}
	return &dataset, nil
}

type AllocSummary struct {
	Category string
	Data     interface{}
	Error    error
}
