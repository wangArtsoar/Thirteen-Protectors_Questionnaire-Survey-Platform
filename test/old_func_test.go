package main

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
)

func Test(t *testing.T) {
	//ipList := generateIPList()
	//fmt.Println(ipList)
	element := removeElement([]int{1, 2, 3, 4}, 5)
	fmt.Println(element)
}

func generateIPList() string {
	ipList := ""
	for i := 0; i < 255; i++ {
		ip := fmt.Sprintf("%d.%d.%d", i, i+1, i+2)
		ipList += ip + ","
	}
	return ipList
}

func removeElement(nums []int, val int) int {
	left := 0
	for _, v := range nums { // v 即 nums[right]
		if v != val {
			nums[left] = v
			left += 1
		}
	}
	return left
}

func Test_2(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	copyArr := make([]int, len(arr))
	copy(copyArr, arr) // 复制原数组

	// 在 copy 后的数组上进行修改
	copyArr[2] = 99

	fmt.Println(arr)     // 输出原数组
	fmt.Println(copyArr) // 输出修改后的数组
}

func Test_3(t *testing.T) {
	// 创建一个切片
	mySlice := []string{"hello", "world", "this", "is", "a", "test"}

	// 向切片中添加元素
	newElement := "append me!"
	appendedSlice := append(mySlice, newElement)

	// 检查数组是否已更改
	fmt.Println("Original array before append:", mySlice)
	fmt.Println("Appended array after append:", appendedSlice)
}

func workerTest_4(wg *sync.WaitGroup, num int) {
	defer wg.Done()
	fmt.Printf("Worker %d is running...\n", num)
}

func Test_4(t *testing.T) {

	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go workerTest_4(&wg, i)
	}

	wg.Wait()
	fmt.Println("All workers have finished.")
}

func workerTest_5(wg *sync.WaitGroup, num int) {
	defer wg.Done()
	fmt.Printf("Worker %d is running...\n", num)
}

func Test_5(t *testing.T) {
	numWorkers := 3
	pool := make(chan int, numWorkers)

	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		go func(workerNum int) {
			workerTest_5(&wg, workerNum)
			wg.Done()
		}(i)
		pool <- i
	}

	go func() {
		wg.Wait()
		close(pool)
	}()
}

func TestMultiply(t *testing.T) {
	s := multiply("123", "123")
	fmt.Println(s)
}

// num1,num2 = "0"
// num1 = "1",return num2
// 123 * 123 = 15129
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	m, n := len(num1), len(num2)
	res := make([]int, m+n)
	for i := m - 1; i >= 0; i-- {
		n1, _ := strconv.Atoi(string(num1[i]))
		for j := n - 1; j >= 0; j-- {
			n2, _ := strconv.Atoi(string(num2[j]))
			mul := n1 * n2
			p1, p2 := i+j, i+j+1
			sum := mul + res[p2]
			res[p2] = sum % 10
			res[p1] += sum / 10
		}
	}
	for len(res) > 1 && res[0] == 0 {
		res = res[1:]
	}
	result := ""
	for _, v := range res {
		result += strconv.Itoa(v)
	}
	return result
}
