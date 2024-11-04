package goroutines

import "fmt"

type (
	// LazyEvaluator is a function that returns a value.
	LazyEvaluator func(any) (any, any)
)

func BuildLazyEvaluator(evalFunc LazyEvaluator, initState any) func() any {
	retValChan := make(chan any)

	loopFunc := func() {
		var actState = initState
		var retValue any
		for {
			retValue, actState = evalFunc(actState)
			retValChan <- retValue
		}
	}

	retFunc := func() any {
		return <-retValChan
	}

	go loopFunc()
	return retFunc
}

func BuildLazyIntEvaluator(evalFunc LazyEvaluator, initState any) func() int {
	ef := BuildLazyEvaluator(evalFunc, initState)
	return func() int {
		return ef().(int)
	}
}

func run_lazy_evaluator() {
	evenFunc := func(state any) (any, any) {
		originalState := state.(int)
		newState := originalState + 2
		return originalState, newState
	}

	even := BuildLazyIntEvaluator(evenFunc, 0)

	for i := 0; i < 10; i++ {
		fmt.Printf("%vth even: %v\n", i, even())
	}
}
