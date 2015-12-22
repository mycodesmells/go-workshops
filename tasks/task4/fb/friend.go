package fb

import "fmt"

// Person is a single instance of people.
type Person struct {
	Name     string
	Location string
}

// Friend in need is a friend indeed.
type Friend interface {
	TellSecret() string
}

// TalkToFriend allows you to pretend you care about other people.
func TalkToFriend(f Friend) string {
	s := f.TellSecret()

	fmt.Println(" - What is the secret, dude?")
	fmt.Printf(" - %s\n", s)

	return s
}
