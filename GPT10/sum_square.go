package GPT10

import (
	"fmt"
	"sync"
)

func square(wg *sync.WaitGroup, input <-chan int, output chan<- int) {
	defer wg.Done()
	for num := range input {
		result := num * num
		output <- result
	}
}

func sum(wg *sync.WaitGroup, input <-chan int, result chan<- int) {
	defer wg.Done()
	total := 0
	for num := range input {
		total += num
	}
	result <- total
}

func SumOfSquare() {
	input := make(chan int)
	squareOutput := make(chan int)
	sumOutput := make(chan int)
	var wg sync.WaitGroup

	// 启动四个协程进行平方计算package GPT10
	//
	//import (
	//	"fmt"
	//	"sync"
	//)
	//
	//func square(wg *sync.WaitGroup, input <-chan int, output chan<- int) {
	//	defer wg.Done()
	//	for num := range input {
	//		result := num * num
	//		output <- result
	//	}
	//}
	//
	//func sum(wg *sync.WaitGroup, input <-chan int, result chan<- int) {
	//	defer wg.Done()
	//	total := 0
	//	for num := range input {
	//		total += num
	//	}
	//	result <- total
	//}
	//
	//func SumOfSquare() {
	//	input := make(chan int)
	//	squareOutput := make(chan int)
	//	sumOutput := make(chan int)
	//	var wg sync.WaitGroup
	//
	//	// 启动四个协程进行平方计算
	//	for i := 0; i < 4; i++ {
	//		wg.Add(1)
	//		go square(&wg, input, squareOutput)
	//	}
	//
	//	// 启动一个协程进行求和计算
	//	wg.Add(1)
	//	go sum(&wg, squareOutput, sumOutput)
	//
	//	// 发送数据到输入通道
	//	for i := 1; i <= 10; i++ {
	//		input <- i
	//	}
	//	close(input)
	//
	//	// 等待所有协程完成
	//	wg.Wait()
	//
	//	// 输出最终结果
	//	result := <-sumOutput
	//	fmt.Println("Result:", result)
	//}
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go square(&wg, input, squareOutput)
	}

	// 启动一个协程进行求和计算
	wg.Add(1)
	go sum(&wg, squareOutput, sumOutput)

	// 发送数据到输入通道
	for i := 1; i <= 10; i++ {
		input <- i
	}
	close(input)

	// 等待所有协程完成
	wg.Wait()

	// 输出最终结果
	result := <-sumOutput
	fmt.Println("Result:", result)
}
