package sel_test

import (
	"testing"

	"fmt"

	"."
)

func TestWaitForToiletSuccess(t *testing.T) {
	toiletFree := make(chan bool)
	bladder := make(chan string)

	go sel.WaitForToilet(toiletFree, bladder)
	toiletFree <- true

	state := <-bladder

	if state != "yes, thank you!!" {
		t.Errorf("You should be able to use the toilet by now. Can you?")
	}
}

func TestWaitForToiletFailure(t *testing.T) {
	toiletFree := make(chan bool)
	bladder := make(chan string)

	go sel.WaitForToilet(toiletFree, bladder)

	state := <-bladder
	fmt.Println(state)
	if state != "eh, nevermind..." {
		t.Errorf("Did you go pee outside like some caveman?!")
	}
}
