package slowestkey

func slowestKey(releaseTimes []int, keysPressed string) byte {
	max, key := releaseTimes[0], keysPressed[0]

	for i := 1; i < len(releaseTimes); i++ {
		releaseTime := releaseTimes[i] - releaseTimes[i-1]
		if releaseTime >= max {
			if releaseTime == max {
				if keysPressed[i] > key {
					key = keysPressed[i]
				}
			} else {
				max = releaseTime
				key = keysPressed[i]
			}
		}
	}

	return key
}
