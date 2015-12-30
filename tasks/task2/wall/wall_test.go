package wall

import (
	"testing"
)

func TestWallStats(t *testing.T) {
	stats := WallStats()

	userId := stats["UserId"]
	frCount := stats["FriendsCount"]
	expUserId := 123
	expFrCount := 0

	if userId != expUserId {
		t.Errorf("UserId is %d, expected: %d", userId, expUserId)
	}
	if frCount != expFrCount {
		t.Errorf("You think you have %d friends, expected: %d", frCount, expFrCount)
	}
}
