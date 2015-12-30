package fb

import (
	"testing"
)

func TestCreateImaginaryFriend(t *testing.T) {
	f := Friend{}

	if f.Name != "Johny Bravo" {
		t.Errorf("I expected you to be friends with Johny Bravo, but you hang out with '%s' instead!", f.Name)
	}
	if f.Location != "Aron City" {
		t.Errorf("They were supposed to be from Aron City! Where is %s, anyway?!", f.Name)
	}
}
