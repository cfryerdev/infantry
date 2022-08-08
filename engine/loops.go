package engine

// https://stackoverflow.com/a/25306241
// https://stackoverflow.com/a/16466581

import (
	"fmt"
	"sync"
	"time"
)

func ExecuteEachSecondConcurrent(maxSeconds int, delegate func()) {
	d := time.NewTicker(1 * time.Second)

	exitChannel := make(chan bool)

	go func() {
		time.Sleep(time.Duration(maxSeconds) * time.Second)
		exitChannel <- true
	}()

	for {
		select {
		case <-exitChannel:
			fmt.Println("Completed!")
			d.Stop()
			return
		case timeElapsed := <-d.C:
			fmt.Printf("elapsed: %+v\n", timeElapsed)
			delegate()
		}
	}
}

// ExecuteNumberOfTimesConcurrent Executes a function n number of times with a max concurrent number using a guard
func ExecuteNumberOfTimesConcurrent(times int, delegate func()) {
	var wg sync.WaitGroup
	wg.Add(times)
	for i := 1; i <= times; i++ {
		go func(n int) {
			delegate()
			wg.Done()
		}(i)
	}
	wg.Wait()
}
