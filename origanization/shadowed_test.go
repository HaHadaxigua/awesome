package origanization

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestShadowed(t *testing.T) {
	createPointerReferInt := func() (*int, error) {
		v := rand.Int()
		return &v, nil
	}

	var (
		x *int
	)
	if rand.Int() > 1 {
		x, err := createPointerReferInt()
		if err != nil {
			t.Error(err)
			return
		}
		fmt.Println(x)
	} else {
		x, err := createPointerReferInt()
		if err != nil {
			t.Error(err)
			return
		}
		fmt.Println(x)
	}
	fmt.Println(x) // x will always be nil, because
}
