package numbers

import (
	"fmt"
	"testing"
)

func TestGetAll(t *testing.T) {

	d, err := GetData()
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Printf("%+v\n", d)
}
