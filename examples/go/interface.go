package _go

type mapInt = interface{ func(int) int }

type Len[I interface{ int }] interface {
	Len() int
}

func X4[T int | interface{ int64 }, M mapInt](t T, m M) {

	var n M = func(i int) int { return i + 1 }
	n(3)
}
