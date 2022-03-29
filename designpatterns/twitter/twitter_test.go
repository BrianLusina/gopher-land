package twitter

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestTwitter(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Twitter Suite")
}

var _ = Describe("Twitter Test Suite", func() {

	It("post tweet should add tweet to user's tweets", func() {
		twitter := NewTwitter()
		userId := 1
		tweetId := 1
		twitter.PostTweet(userId, tweetId)
		newsFeed := twitter.GetNewsFeed(userId)
		expected := []int{tweetId}
		Expect(newsFeed).To(Equal(expected))
	})

	It("follow should add user to user's followers", func() {
		twitter := NewTwitter()
		userId := 1
		followeeId := 2
		twitter.Follow(userId, followeeId)

		// let the followee post a tweet
		tweetId := 2
		twitter.PostTweet(followeeId, tweetId)

		// let the user get the news feed
		newsFeed := twitter.GetNewsFeed(userId)
		expected := []int{tweetId}

		Expect(newsFeed).To(Equal(expected))
	})

	It("unfollow should remove user from user's following", func() {
		twitter := NewTwitter()

		followerId := 1
		followeeId := 2

		// let the followee & follower post tweets
		followeeTweetId := 2
		followerTweetId := 3

		// let the user follow the followee
		twitter.Follow(followerId, followeeId)

		// both users post tweets
		twitter.PostTweet(followerId, followerTweetId)
		twitter.PostTweet(followeeId, followeeTweetId)

		// let the follower get their news feed
		actualFollowerNewsFeed := twitter.GetNewsFeed(followerId)

		expectedFollowerNewsFeed := []int{followeeTweetId, followerTweetId}
		Expect(actualFollowerNewsFeed).To(Equal(expectedFollowerNewsFeed))

		// let the followee & follower post new tweets
		followeeTweet2Id := 3
		followerTweet2Id := 4

		twitter.Unfollow(followerId, followeeId)
		twitter.PostTweet(followeeId, followeeTweet2Id)
		twitter.PostTweet(followerId, followerTweet2Id)

		// let the follower get their news feed
		actualFollowerNewsFeed2 := twitter.GetNewsFeed(followerId)
		expectedFollowerNewsFeed2 := []int{followerTweet2Id, followerTweetId}

		// the user's news feed should not show the followee's tweet
		Expect(actualFollowerNewsFeed2).To(Equal(expectedFollowerNewsFeed2))
	})

	It("postTweet(1,5) -> getNewsFeed(1) -> follow(1,2) -> postTweet(2,6) -> getNewsFeed(1) -> unfollow(1,2) -> getNewsFeed(1)", func() {
		twitter := NewTwitter()

		followerId := 1
		followeeId := 2

		followerFirstTweet := 5

		// follower posts their first tweet
		twitter.PostTweet(followerId, followerFirstTweet)
		actualFollowerFirstNewsFeed := twitter.GetNewsFeed(followerId)

		expectedFollowerFirstNewsFeed := []int{followerFirstTweet}
		Expect(actualFollowerFirstNewsFeed).To(Equal(expectedFollowerFirstNewsFeed))

		// follower follows the followee
		twitter.Follow(followerId, followeeId)

		// followee posts a tweet
		followeeFirstTweet := 6
		twitter.PostTweet(followeeId, followeeFirstTweet)

		actualFollowerSecondNewsFeed := twitter.GetNewsFeed(followerId)
		expectedFollowerSecondNewsFeed := []int{followeeFirstTweet, followerFirstTweet}
		Expect(actualFollowerSecondNewsFeed).To(Equal(expectedFollowerSecondNewsFeed))

		// follower unfollows followee
		twitter.Unfollow(followerId, followeeId)
		actualFollowerThirdNewsFeed := twitter.GetNewsFeed(followerId)

		expectedFollowerThirdNewsFeed := []int{followerFirstTweet}
		Expect(actualFollowerThirdNewsFeed).To(Equal(expectedFollowerThirdNewsFeed))
	})

	It("should return 10 most recent tweets in news feed in correct order", func() {
		twitter := NewTwitter()

		userId := 1

		tweets := []int{5, 3, 101, 13, 10, 2, 94, 505, 333, 22, 11}

		for _, tweet := range tweets {
			twitter.PostTweet(userId, tweet)
		}

		actualNewsFeed := twitter.GetNewsFeed(userId)
		expectedNewsFeed := []int{11, 22, 333, 505, 94, 2, 10, 13, 101, 3}

		Expect(actualNewsFeed).To(Equal(expectedNewsFeed))
	})

})
