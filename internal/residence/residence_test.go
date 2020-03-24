package residence

import (
	"fmt"
	"testing"
)

func TestGetData(t *testing.T) {
	r, err := GetData()

	if err != nil {
		t.Error(err)
		return
	}

	fmt.Printf("%+v\n", r)
}
