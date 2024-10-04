package semaphore

import (
	"fmt"
	"sync"
	"time"
)

type Semaphore struct {
	capacity int
	buffer   chan struct{}
}

func New(capacity int) *Semaphore {
	return &Semaphore{
		capacity: capacity,
		buffer:   make(chan struct{}, capacity),
	}
}

func (s *Semaphore) Acquire() {
	s.buffer <- struct{}{}
}

func (s *Semaphore) Release() {
	<-s.buffer
}

func RunSemaphoreExample(capacity int) {
	semaphore := New(capacity)
	wg := &sync.WaitGroup{}

	for i := 0; i < 20; i++ {
		wg.Add(1)

		go func(goroutineNumber int) {
			defer wg.Done()

			semaphore.Acquire()
			defer semaphore.Release()

			fmt.Printf("Goroutine %d is running a task\n", goroutineNumber)
			time.Sleep(2 * time.Second)
		}(i)
	}

	wg.Wait()
	fmt.Println("All goroutines have finished!")
}
