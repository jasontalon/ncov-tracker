package hospital

import (
	"fmt"
	"testing"
)

func TestGetData(t *testing.T) {
	d, err := GetData()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("%+v", d)
}
