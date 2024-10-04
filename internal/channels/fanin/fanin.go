package fanin

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
)

// JoinChannels реализует паттерн FanIn
func JoinChannels[T any](channels ...<-chan T) <-chan T {
	joinedChannel := make(chan T)

	go func() {
		wg := sync.WaitGroup{}

		for _, channel := range channels {
			wg.Add(1)

			go func(ch <-chan T) {
				for value := range ch {
					joinedChannel <- value
				}

				wg.Done()
			}(channel)
		}

		go func() {
			wg.Wait()
			close(joinedChannel)
		}()
	}()

	return joinedChannel
}

const (
	channelsNumber          int = 3
	numbersToWriteInChannel int = 3
)

func RunJoinChannels() {
	channelsToJoin := make([]<-chan int, 0, channelsNumber)
	wg := sync.WaitGroup{}

	for i := 0; i < channelsNumber; i++ {
		wg.Add(1)

		go func() {
			channel := make(chan int)
			channelsToJoin = append(channelsToJoin, channel)
			wg.Done()

			for j := 0; j < numbersToWriteInChannel; j++ {
				channel <- (rand.Intn(10) + 1) * int(math.Pow(10, float64(i)))
			}

			close(channel)
		}()
	}

	wg.Wait()
	for value := range JoinChannels(channelsToJoin...) {
		fmt.Println(value)
	}
}
