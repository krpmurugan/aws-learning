package main

import "fmt"

type RoundRobin struct {
	servers []string
	counter int
}

func (rr *RoundRobin) AddServer(server string) {
	rr.servers = append(rr.servers, server)
}

func (rr *RoundRobin) NextServer() string {
	if len(rr.servers) == 0 {
		return ""
	}
	server := rr.servers[rr.counter]
	rr.counter = (rr.counter + 1) % len(rr.servers)
	return server
}

func main() {
	rr := &RoundRobin{}

	rr.AddServer("server1")
	rr.AddServer("server2")
	rr.AddServer("server3")

	for i := 0; i < 5; i++ {
		server := rr.NextServer()
		fmt.Println("Request sent to:", server)
	}
}
