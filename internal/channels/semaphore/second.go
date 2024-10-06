package semaphore

import (
	"fmt"
	"sync"
	"time"
)

func Foo() {
	wg := sync.WaitGroup{}
	semaphore := New(3)

	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			semaphore.Acquire()
			defer semaphore.Release()
			printNumber(i)
		}()
	}

	wg.Wait()
}

func printNumber(n int) {
	time.Sleep(time.Second)
	fmt.Println(n)
}
