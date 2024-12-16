package channels

import (
	"fmt"
	"sync"
)

// sumChannels
// in - слайс входных каналов, в которые приходят числа
// Признак окончания данных в канале - канал закрыт
func sumChannels(inputs []chan int64) int64 {
	// ваш код
	wg := sync.WaitGroup{}
	sumChannel := make(chan int64, len(inputs))
	result := make(chan int64)
	for _, c := range inputs {
		wg.Add(1)
		go func(inChannel chan int64, wg *sync.WaitGroup) {
			var sum int64
			for {
				res, ok := <-inChannel
				if !ok {
					fmt.Println("Failed to read from imChannel in goroutine")
					sumChannel <- sum
					wg.Done()
					break
				}
				sum += res
			}
		}(c, &wg)
	}
	wg.Wait()
	close(sumChannel)
	go func(sums <-chan int64, res *chan int64) {
		var sum int64
		//resChannel := make(chan int64)
		for {
			r, ok := <-sums
			if !ok {
				fmt.Printf("result: %v\n", sum)
				*res <- sum
				break
			}
			sum += r

		}
	}(sumChannel, &result)
	vv := <-result
	fmt.Printf("return: %v\n", vv)
	return vv
}

