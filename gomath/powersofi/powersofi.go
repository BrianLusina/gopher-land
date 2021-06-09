package powersofi

func PowersOfI(n int) string {
	vals := []string{"1", "i", "-1", "-i"}
	return vals[n%4]
}
