package engine

// https://stackoverflow.com/a/25306241
// https://stackoverflow.com/a/16466581

import (
	"sync"
	"time"
)

// ExecuteEachSecond Executes a function each second
func ExecuteEachSecond(seconds int, delegate func()) {
	for _ = range time.Tick(time.Duration(seconds) * time.Second) {
		delegate()
	}
}

// ExecuteEachSecondConcurrent Executes a function each second concurrently
func ExecuteEachSecondConcurrent(seconds int, delegate func()) {
	var wg sync.WaitGroup
	wg.Add(seconds)
	for _ = range time.Tick(time.Duration(seconds) * time.Second) {
		go func() {
			delegate()
			wg.Done()
		}()
	}
	wg.Wait()
}

// ExecuteNumberOfTimes Executes a function n number of times
func ExecuteNumberOfTimes(times int, delegate func()) {
	for i := 1; i <= times; i++ {
		delegate()
	}
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
			<-guard
			wg.Done()
		}(i)
	}
	wg.Wait()
}
