package user

type UserIterator struct {
	index int
	users []*User
}

func (i *UserIterator) hasNext() bool {
	if i.index < len(i.users) {
		return true
	}
	return false
}

func (i *UserIterator) next() *User {
	if i.hasNext() {
		user := i.users[i.index]
		i.index++
		return user
	}
	return nil
}
