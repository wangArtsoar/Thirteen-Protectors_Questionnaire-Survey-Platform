package test_0804

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func Test04ClaudeAnswer(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			worker(num)
		}(i)
	}
	wg.Wait()
	fmt.Println("All workers have finished.")
}

func Test04FirstAnswer(t *testing.T) {
	numWorkers := 3
	pool := make(chan int, numWorkers)
	var wg sync.WaitGroup
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go func(workerNum int) {
			worker(workerNum)
			wg.Done()
		}(i)
		pool <- i
	}
	go func() {
		wg.Wait()
		close(pool)
	}()
	for range pool {
		fmt.Println("All workers have finished.")
	}
}

func Test04SecondAnswer(t *testing.T) {
	const processes = 3
	var wg sync.WaitGroup
	pool := make(chan int, processes)
	for i := 0; i < processes; i++ {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			worker(num)
		}(i)
	}
	go func() {
		wg.Wait()
		close(pool)
	}()
	for i := 0; i < 5; i++ {
		pool <- i
	}
	time.Sleep(time.Second)
	fmt.Println("All workers have finished.")
}

func Test04ThirdAnswer(t *testing.T) {
	var wg sync.WaitGroup
	pool := make(chan int, 3)
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(num int) {
			worker(num)
			wg.Done()
		}(i)
	}
	go func() {
		wg.Wait()
		close(pool)
	}()
	for i := 0; i < 5; i++ {
		pool <- i
	}
	close(pool)
	for result := range pool {
		fmt.Printf("Worker %d is running...\n", result)
	}
}

func Test04FourthAnswer(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(3)

	pool := make(chan int, 3)

	for i := 0; i < 5; i++ {
		go func(num int) {
			worker(num)
			wg.Done()
		}(i)
		pool <- i
	}

	close(pool)
	wg.Wait()
	fmt.Println("All workers have finished.")
}

func Test04FifthAnswer(t *testing.T) {
	var wg sync.WaitGroup
	poolSize := 3
	pool := make(chan int, poolSize)

	for i := 0; i < poolSize; i++ {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			worker(num)
		}(i)
	}

	for i := 0; i < 5; i++ {
		pool <- i
	}
	close(pool)

	wg.Wait()
	fmt.Println("All workers have finished.")
}

func worker(num int) {
	fmt.Printf("Worker %d is running...\n", num)
}
