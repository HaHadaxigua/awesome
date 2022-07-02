package main

type constraint[T any] interface {
	func(...T) T
}

func FN[A any, F constraint[A]](v F, args ...A) {
	v(args...)
}

func main() {
	fnInt := func(v ...int) int {
		r := v[0] * v[0]
		println(r)
		return r
	}

	fnFloat := func(v ...float32) float32 {
		var r float32
		for _, e := range v {
			r += e
		}
		println(r)
		return r
	}

	FN(fnInt, 2)
	FN(fnFloat, 2, 3, 4)
}
