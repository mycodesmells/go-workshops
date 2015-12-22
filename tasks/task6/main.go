package main

import (
	"fmt"
	//	"sync"

	"github.com/slomek/go-workshops/tasks/task6/dogs"
)

func main() {
	dogs.WhoLetThemOut()

	// Two lines here

	// Add something to each call
	go dogs.Who()
	go dogs.Who()
	go dogs.Who()

	// One line here

	fmt.Println()
	fmt.Println("There should be at least three `Who?`'s at this point.")
}
