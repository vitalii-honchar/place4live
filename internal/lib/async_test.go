package lib

import "testing"

func TestAsync(t *testing.T) {
	// GIVEN
	v := 10

	// WHEN
	actual := <-Async(func() int {
		return v
	})

	// THEN
	if actual != v {
		t.Errorf("Unexpected return value from Async: expected = %v, actual = %v\n", v, actual)
	}
}
