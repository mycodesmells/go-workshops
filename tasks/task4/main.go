package main

import (
	"fmt"

	"github.com/slomek/go-workshops/tasks/task4/fb"
)

// SuperHero is kind of hero, but is super at the same time.
type SuperHero struct {
	fb.Person
	HeroName string
}

// TellSecret provides you with some information you did not know earlier.
func (b SuperHero) TellSecret() string {
	return fmt.Sprintf("I'm %s!", b.HeroName)
}

func main() {
	fmt.Println("Is it true? Another friend of yours?")

	sh := SuperHero{
		Person:   fb.Person{Name: "Bruce Wayne", Location: "Gotham City"},
		HeroName: "Batman"}

	fb.TalkToFriend(sh)
}
