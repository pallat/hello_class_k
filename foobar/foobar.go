package foobar

import "strconv"

func FooBar(n int) string {
	if n == 3 {
		return "Foo"
	}

	return strconv.Itoa(n)
}
