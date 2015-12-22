package bday

// FakeFriend is somebody you hang out with because you have nowhere else to go.
type FakeFriend struct {
	Age int
}

// HappyBirthday is a representation of happiness of people who, unless B-day girl/boy, are not getting older that day.
func (b FakeFriend) HappyBirthday() {
	b.Age++
}
