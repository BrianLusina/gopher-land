package gigasecond

import "time"

const gigasecond = 1e9

// AddGigasecond takes in a moment and returns a gigasecond after that moment
func AddGigasecond(moment time.Time) time.Time {
	return moment.Add(gigasecond * time.Second)
}
