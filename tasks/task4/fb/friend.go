package fb
import "fmt"

type Person struct {
	Name string
	Location string
}

type Friend interface {
	TellSecret() string
}

func TalkToFriend(f Friend) string {
	s := f.TellSecret()

	fmt.Println(" - What is the secret, dude?")
	fmt.Printf(" - %s\n", s)

	return s
}
