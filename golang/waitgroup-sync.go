package main

import (
	"fmt"
	"sync"
)

func Sum(arr []int, wg *sync.WaitGroup, m *sync.Mutex, ch chan int) {
	defer wg.Done()
	sum := 0
	for _, v := range arr {
		sum += v
	}
	m.Lock()
	ch <- sum
	m.Unlock()
}

func main() {
	var wg sync.WaitGroup
	var m sync.Mutex
	ch := make(chan int, 0)
	arr := []int{5, 5}
	wg.Add(3)
	go Sum(arr, &wg, &m, ch)
	a := <-ch
	go Sum(arr, &wg, &m, ch)
	b := <-ch
	arrq := []int{10, 20, 30}
	go Sum(arrq, &wg, &m, ch)
	c := <-ch
	wg.Wait()
	fmt.Println(a + b + c) //o/p: 80
}
