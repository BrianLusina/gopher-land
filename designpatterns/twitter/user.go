package twitter

// user represents a Twitter user.
// id is the unique identifier for the user.
// tweets is a slice of the user's tweet ids
type (
	userid = int
	user   struct {
		id        userid
		tweets    []Tweet
		followers map[userid]*user
		following map[userid]*user
	}
)

func newUser(userId int) *user {
	return &user{
		id:        userId,
		tweets:    []Tweet{},
		followers: make(map[userid]*user),
		following: make(map[userid]*user),
	}
}

func (u *user) getFollowers() []int {
	var followers []int
	for _, follower := range u.followers {
		followers = append(followers, follower.id)
	}
	return followers
}

func (u *user) addFollower(follower *user) {
	u.followers[follower.id] = follower
}

func (u *user) removeFollowing(followingId int) {
	delete(u.following, followingId)
}
func (u *user) removeFollower(followerId int) {
	delete(u.followers, followerId)
}

func (u *user) addFollowing(followee *user) {
	u.following[followee.id] = followee
}

func (u *user) getFollowing() []*user {
	var following []*user
	for _, followee := range u.following {
		following = append(following, followee)
	}
	return following
}

func (u *user) getTweets() []Tweet {
	return u.tweets
}
