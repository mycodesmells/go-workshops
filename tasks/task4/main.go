package main

import (
	"fmt"

	"github.com/slomek/go-workshops/tasks/task4/fb"
)

type SuperHero struct {
	fb.Person
	HeroName string
}

func (b SuperHero) TellSecret() string {
	return fmt.Sprintf("I'm %s!", b.HeroName)
}

func main() {
	fmt.Println("Is it true? Another friend of yours?")

	sh := SuperHero{
		Person: fb.Person{Name: "Bruce Wayne",Location: "Gotham City"},
		HeroName: "Batman"}

	fb.TalkToFriend(sh)
}
