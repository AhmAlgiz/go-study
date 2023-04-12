package observer

type User struct {
	Observer
	data string
}

func (u *User) Update(data string) {
	u.data = data
}
