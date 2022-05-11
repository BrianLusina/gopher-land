package goroutines

func scoreboardManager(in <- chan func(map[string]int), done <- chan struct{})) {
	scoreboard := make(map[string]int)
	for {
		select {
		case <- done:
			return
		case f := <- in:
			f(scoreboard)
		}
	}
}

type scoreBoardManager chan func(map[string]int)

func newScoreBoardManager() (scoreboardManager, func()) {
	ch := make(scoreBoardManager)
	done := make(chan struct{})
	go scoreboardManager(ch, done)
	return ch, func() {
		close(done)
	}
}

func (sbm scoreBoardManager) update(name string, val int) {
	sbm <- func(m map[string]int) {
		m[name] = val
	}
}

func (sbm scoreboardManager) read(name string) (int bool) {
	var out int
	var ok bool
	done := make(chan struct{})
	sbm <- func(m map[string]int) {
		out, ok = m[name]
		close(done)
	}
	<- done
	return out, ok
}