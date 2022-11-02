package foobar

import "testing"

/*
!	given 1 wants "1"
!	given 2 wants "2"
!	given 3 wants "Foo"
!	given 4 wants "4"
!	given 5 wants "Bar"
!	given 6 wants "Foo"
*/

func TestFooBarGiven1(t *testing.T) {
	given := 1
	want := "1"

	result := FooBar(given)

	if want != result {
		t.Errorf("given %d wants %q but %q", given, want, result)
	}
}
func TestFooBarGiven2(t *testing.T) {
	given := 2
	want := "2"

	result := FooBar(given)

	if want != result {
		t.Errorf("given %d wants %q but %q", given, want, result)
	}
}
func TestFooBarGiven3(t *testing.T) {
	given := 3
	want := "Foo"

	result := FooBar(given)

	if want != result {
		t.Errorf("given %d wants %q but %q", given, want, result)
	}
}
func TestFooBarGiven4(t *testing.T) {
	given := 4
	want := "4"

	result := FooBar(given)

	if want != result {
		t.Errorf("given %d wants %q but %q", given, want, result)
	}
}
