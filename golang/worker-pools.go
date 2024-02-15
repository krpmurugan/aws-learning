package main

import (
	"fmt"
	"sync"
)

func Workers(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d is processing job %d\n", id, job)
	}
}

func main() {
	numWorkers := 3
	numJobs := 10

	jobs := make(chan int, numJobs)

	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go Workers(i, jobs, &wg)
	}

	for i := 0; i < numJobs; i++ {
		jobs <- i
	}

	close(jobs)

	wg.Wait()

}
