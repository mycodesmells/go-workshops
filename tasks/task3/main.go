package main

import (
	"fmt"

	"github.com/slomek/go-workshops/tasks/task3/fb"
)

func main() {
	md := fb.Friend{"Matt Damon", "Mars"}

	fmt.Println("OMG You have a friend after all:")
	fmt.Printf(" - Hello, %s! How's life on %s?\n", md.Name, md.Location)
}
