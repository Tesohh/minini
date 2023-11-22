package data

type User struct {
	Username string
}

func (u User) IsEmpty() bool {
	return u == User{}
}
