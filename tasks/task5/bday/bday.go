package bday

type FakeFriend struct {
	Age int
}

func (b FakeFriend) HappyBirthday() {
	b.Age += 1
}
