package fb_test

import (
	"testing"

	"."
)

type FakeFriend struct {
}

func TestHavingLoyalFriend(t *testing.T) {
	ff := FakeFriend{}
	s := fb.TalkToFriend(ff)

	if s != "You're fat!" {
		t.Errorf("I thought you were a fake friend that would say `You're fat!`, but you said `%s`, instead!", s)
	}
}