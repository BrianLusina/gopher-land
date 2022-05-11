package goroutines

import "sync"

type mutexScoreBoardManager struct {
	mu         sync.RWMutex
	scoreboard map[string]int
}

func newMutexScoreBoardManager() *mutexScoreBoardManager {
	return &mutexScoreBoardManager{
		scoreboard: make(map[string]int),
	}
}

func (msm *mutexScoreBoardManager) update(name string, val int) {
	msm.mu.Lock()
	defer msm.mu.Unlock()
	msm.scoreboard[name] = val
}

func (msm *mutexScoreBoardManager) read(name string) (int bool) {
	msm.mu.RLock()
	defer msm.mu.RUnlock()
	val, ok := msm.scoreboard[name]
	return val, ok
}
