package foobar

import (
	"strconv"
	"time"
)

type Intner interface {
	Intn(n int) int
}

type timeNow interface {
	Now() time.Time
}

func RandomFooBar(r1 Intner) string {
	return FooBar(r1.Intn(4) + 1)
}

func FooBar(n int) string {
	if n == 3 {
		return "Foo"
	}

	return strconv.Itoa(n)
}
