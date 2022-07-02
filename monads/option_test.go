package monads

import (
	"testing"

	"github.com/samber/mo"
	"gotest.tools/assert"
)

func Test(t *testing.T) {
	option1 := mo.Some(42)
	// Some(42)
	assert.Equal(t, option1.MustGet(), 42)

	option1.
		FlatMap(func(value int) mo.Option[int] {
			return mo.Some(value * 2)
		}).
		FlatMap(func(value int) mo.Option[int] {
			return mo.Some(value % 2)
		}).
		FlatMap(func(value int) mo.Option[int] {
			return mo.Some(value + 21)
		}).
		OrElse(1234)
	// 21

	option2 := mo.None[int]()
	// None

	option2.OrElse(1234)
	// 1234

	option3 := option1.Match(
		func(i int) (int, bool) {
			// when value is present
			return i * 2, true
		},
		func() (int, bool) {
			// when value is absent
			return 0, false
		},
	)
	println(option3.Get())
}
