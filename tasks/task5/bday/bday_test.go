package bday

import (
	"testing"
)

func TestFriendGettingOlder(t *testing.T) {
	ff := FakeFriend{Age: 77}
	ff.HappyBirthday()

	if ff.Age != 78 {
		t.Errorf("Your friend is %d years old, but should be: %d", ff.Age, 78)
	}
}
