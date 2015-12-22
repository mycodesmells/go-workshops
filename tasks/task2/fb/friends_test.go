package fb_test

import (
	"testing"

	"."
)

func TestNumberOfFriends(t *testing.T) {
	expected := 1
	n := fb.NumberOfFriends()

	if n != expected {
		t.Errorf("You think you have %d friends, expected: %d", n, expected)
	}
}
