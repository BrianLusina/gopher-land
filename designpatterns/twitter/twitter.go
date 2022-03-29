// Package twitter is a Go package that implements the Twitter design pattern.
package twitter

import (
	"sort"
)

// Twitter implementation
type Twitter struct {
	users map[int]*user
}

func NewTwitter() Twitter {
	return Twitter{
		users: map[int]*user{},
	}
}

// PostTweet posts a tweet and adds the tweet to the user's tweets
// in this contrived example, we assume the user already exists.
// Therefore, we need to 'create' a user or add them to the system.
func (t *Twitter) PostTweet(userId int, tweetId int) {
	user, ok := t.users[userId]

	if !ok {
		user = t.addUser(userId)
	}
	user.tweets = append(user.tweets, newTweet(tweetId))
}

// GetNewsFeed gets the news feed for a user, retrieving the 10 most recent
// Each item in the news feed must be posted by users who the user followed or by the user themselves.
// Tweets must be ordered from most recent to least recent.
func (t *Twitter) GetNewsFeed(userId int) []int {
	if user, ok := t.users[userId]; ok {
		tweets := user.getTweets()
		following := user.getFollowing()

		for _, followingUser := range following {
			tweets = append(tweets, followingUser.getTweets()...)
		}

		// sort tweets by timestamp
		sort.Sort(Tweets(tweets))

		// get most recent 10
		if len(tweets) > 10 {
			tweets = tweets[:10]
		}
		newsFeed := make([]int, 0, 10)
		for _, tweet := range tweets {
			newsFeed = append(newsFeed, tweet.id)
		}
		return newsFeed
	}
	return []int{}
}

// Follow follows a user.
// followerId the user following
// followeeId the user being followed
func (t *Twitter) Follow(followerId int, followeeId int) {
	follower, ok := t.users[followerId]

	if !ok {
		follower = t.addUser(followerId)
	}

	followee, ok := t.users[followeeId]
	if !ok {
		followee = t.addUser(followeeId)
	}
	followee.addFollower(follower)
	follower.addFollowing(followee)
}

func (t *Twitter) Unfollow(followerId int, followeeId int) {
	follower, ok := t.users[followerId]

	if !ok {
		return
	}

	followee, ok := t.users[followeeId]
	if !ok {
		return
	}

	follower.removeFollowing(followee.id)
	followee.removeFollower(follower.id)
}

func (t *Twitter) addUser(userId int) *user {
	user := newUser(userId)
	t.users[userId] = user
	return user
}
