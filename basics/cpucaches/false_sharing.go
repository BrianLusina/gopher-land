package cpucaches

import "sync"

type Input struct {
	a int64
	b int64
}

type Result struct {
	sumA int64
	sumB int64
}

type ResultWithPadding struct {
	sumA int64
	// _ extra padding
	_    [56]byte
	sumB int64
}

func countWithoutPadding(inputs []Input) Result {
	wg := sync.WaitGroup{}
	wg.Add(2)

	result := Result{}

	go func() {
		for i := 0; i < len(inputs); i++ {
			result.sumA += inputs[i].a
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < len(inputs); i++ {
			result.sumB += inputs[i].b
		}
		wg.Done()
	}()

	wg.Wait()
	return result
}

func countWithPadding(inputs []Input) ResultWithPadding {
	wg := sync.WaitGroup{}
	wg.Add(2)

	result := ResultWithPadding{}

	go func() {
		for i := 0; i < len(inputs); i++ {
			result.sumA += inputs[i].a
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < len(inputs); i++ {
			result.sumB += inputs[i].b
		}
		wg.Done()
	}()

	wg.Wait()
	return result
}
