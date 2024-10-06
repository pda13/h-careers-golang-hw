package main

import (
	"fmt"
	"sync"
	"time"
)

type Worker struct {
	ID int
}

func NewWorker(id int) *Worker {
	return &Worker{
		ID: id,
	}
}

// Fanout
func (w *Worker) Start(jobsChannel <-chan Job, resultsChannel chan<- int) {
	for job := range jobsChannel {
		fmt.Printf("woker #%d started job #%+v!\n", w.ID, job)
		resultsChannel <- Factorial(job.value)
		time.Sleep(5 * time.Second)
		fmt.Printf("woker #%d finished job #%+v!\n", w.ID, job)
	}
}

func Factorial(n int) int {
	if n <= 1 {
		return 1
	}

	return n * Factorial(n-1)
}

type Job struct {
	value int
}

func NewJob(value int) *Job {
	return &Job{
		value: value,
	}
}

const (
	JobsNumber    int = 5
	WorkersNumber int = 3
)

func main() {
	jobsChannel := make(chan Job, JobsNumber)
	resultsChannel := make(chan int, JobsNumber)

	//!!!!
	go func() {
		for j := 0; j < JobsNumber; j++ {
			jobsChannel <- *NewJob(j + 1)
		}
		close(jobsChannel)
	}()

	///!!!!
	go func() {
		wg := sync.WaitGroup{}

		for w := 0; w < WorkersNumber; w++ {
			worker := NewWorker(w + 1)

			wg.Add(1)
			go func() {
				defer wg.Done()
				worker.Start(jobsChannel, resultsChannel)
			}()
		}

		wg.Wait()
		close(resultsChannel)
	}()

	for result := range resultsChannel {
		fmt.Println("Result:", result)
	}
}
