package main

import (
	"fmt"

	"github.com/slomek/go-workshops/tasks/task2/fb"
)

func main() {
	fmt.Println("Hello FB!")
	fmt.Printf("I see I have %d friends today.\n", fb.NumberOfFriends)
	fmt.Println("It does not look so good...")
}
