package individual

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	r, err := GetData()

	if err != nil {
		t.Error(err)
		return
	}

	fmt.Printf("%+v\n", r)
}
