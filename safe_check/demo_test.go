package safe_check

import (
	"fmt"
	"testing"
	"time"
)

func TestSafe(t *testing.T) {
	var x int
	duration := time.Duration(time.Duration(x) * time.Second)
	fmt.Println(duration.String())
}
