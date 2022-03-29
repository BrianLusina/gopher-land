package twitter

import "time"

type Tweet struct {
	id        int
	timestamp time.Time
}

type Tweets []Tweet

func (tweets Tweets) Len() int           { return len(tweets) }
func (tweets Tweets) Swap(i, j int)      { tweets[i], tweets[j] = tweets[j], tweets[i] }
func (tweets Tweets) Less(i, j int) bool { return tweets[i].timestamp.After(tweets[j].timestamp) }

func newTweet(id int) Tweet {
	return Tweet{
		id:        id,
		timestamp: time.Now(),
	}
}
