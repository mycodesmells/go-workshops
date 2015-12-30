package wall

import "github.com/slomek/go-workshops/tasks/task2/friends"

func WallStats() map[string]int {
	return map[string]int {
		"UserId": 123,
		"FriendsCount": friends.NumberOfFriends(),
	}
}
