package test_0810

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	fmt.Printf("Worker %d is running...\n", id)

	// 模拟耗时任务
	time.Sleep(time.Second * time.Duration(rand.Intn(10)))
	fmt.Printf("Worker %d finished.\n", id)
	wg.Done()
}

func TestAnswerSecond(t *testing.T) {
	rand.NewSource(time.Now().UnixNano())
	//rand.Seed(time.Now().UnixNano())
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		//wg.Add(2)
		wg.Add(1)
		go func(id int) {
			//worker(id,&wg,&in,&out)
			worker(id, &wg)
		}(i)
	}
	wg.Wait()
}
