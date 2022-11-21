package user

type UserCollection struct {
	users []*User
}

func (u *UserCollection) createIterator() Iterator[User] {
	return &UserIterator{
		users: u.users,
	}
}
