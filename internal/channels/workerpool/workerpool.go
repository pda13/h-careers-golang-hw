package workerpool

import (
	"fmt"
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
	JobsNumber    int = 10
	WorkersNumber int = 3
)

func RunWorkerPoolExample() {
	jobsChannel := make(chan Job, JobsNumber)
	resultsChannel := make(chan int, JobsNumber)

	for w := 0; w < WorkersNumber; w++ {
		worker := NewWorker(w + 1)
		go worker.Start(jobsChannel, resultsChannel)
	}

	for j := 0; j < JobsNumber; j++ {
		jobsChannel <- *NewJob(j + 1)
	}
	close(jobsChannel)

	for r := 0; r < JobsNumber; r++ {
		fmt.Println("Result:", <-resultsChannel)
	}
}
