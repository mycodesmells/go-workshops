package unbuff

import (
	"testing"
)

func toiletKnockKnock(throughDoorTalk chan string) {
}

func TestWaitForToilet(t *testing.T) {
	desparateTalk := make(chan string)

	go toiletKnockKnock(desparateTalk)

	answer := <-desparateTalk
	if answer != "Flushing! Out in a sec!" {
		t.Errorf("Oh come on! My bladder is going to explode here!")
	}
}
