package fb

type Person struct {
	Name string
	Location string
}

type NaivePerson interface {
	LikesMe() bool
	ForceLikeMe()
}

func (np NaivePerson) LikesMe() {
	return np.ableToLike
}
