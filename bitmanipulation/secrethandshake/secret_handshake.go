package secret

var signals = []string{"wink", "double blink", "close your eyes", "jump"}

func Handshake(code uint) []string {
	handshake := []string{}
	switch {
	case code < 1:
	case code&16 == 0:
		for _, signal := range signals {
			if code&1 != 0 {
				handshake = append(handshake, signal)
			}
			code >>= 1
		}
	default:
		for i := 3; i >= 0; i-- {
			if code&8 != 0 {
				handshake = append(handshake, signals[i])
			}
			code <<= 1
		}
	}
	return handshake
}
