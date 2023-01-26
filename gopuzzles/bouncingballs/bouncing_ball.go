package bouncingballs

func bouncingBall(h, bounce, window float64) int {
	rebounds := -1

	if 0 < bounce && bounce < 1 {
		for h > window {
			h *= bounce
			rebounds += 2
		}
	}

	return rebounds
}
