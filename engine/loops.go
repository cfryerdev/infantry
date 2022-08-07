package engine

// https://stackoverflow.com/a/25306241
// https://stackoverflow.com/a/16466581

import (
	"sync"
	"time"
)

// ExecuteEachSecondConcurrent Executes a function each second concurrently with a max count
func ExecuteEachSecondConcurrent(maxSeconds int, delegate func()) {
	var wg sync.WaitGroup
	wg.Add(maxSeconds)
	for _ = range time.Tick(time.Duration(1) * time.Second) {
		go func() {
			delegate()
			wg.Done()
		}()
	}
	wg.Wait()
}

// ExecuteNumberOfTimesConcurrent Executes a function n number of times with a max concurrent number using a guard
func ExecuteNumberOfTimesConcurrent(times int, maxConcurrently int, delegate func()) {
	var guard = make(chan struct{}, maxConcurrently)
	var wg sync.WaitGroup
	wg.Add(times)
	for i := 1; i <= times; i++ {
		go func(n int) {
			guard <- struct{}{}
			delegate()
			wg.Done()
			<-guard
		}(i)
	}
	wg.Wait()
}
