package main

import (
	"fmt"
	"sync"
)

func main() {
	var a sync.WaitGroup
	var b sync.WaitGroup
	var c sync.WaitGroup

	waitChan := make(chan bool)
	a.Add(1)
	b.Add(1)
	c.Add(1)

	go func() {
		var i int = 1

		for j := 0; j < 3; j++ {
			a.Wait()
			a.Add(1)
			fmt.Printf("%d\n", i)
			//time.Sleep(time.Duration(1) * time.Second)
			b.Done()

			i += 3
		}
	}()

	go func() {
		var i int = 2
		for j := 0; j < 3; j++ {
			b.Wait()
			b.Add(1)
			fmt.Printf("%d\n", i)
			//time.Sleep(time.Duration(1) * time.Second)
			c.Done()

			i += 3
		}
	}()

	go func() {
		var i int = 3
		for j := 0; j < 3; j++ {
			c.Wait()
			c.Add(1)
			fmt.Printf("%d\n", i)
			//time.Sleep(time.Duration(1) * time.Second)
			a.Done()

			i += 3
		}
	}()
	a.Done() // Kick things off !
	<-waitChan
}
