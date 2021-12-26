package main

import (
	"fmt"
	"globby/cache"
	"math/rand"
	"time"
)

func producer(lifecycle chan bool) {
	fmt.Println("Producing...")
	defer close(lifecycle)

	connection := cache.Client
	rand.Seed(time.Now().Unix())

	for i := 0; i <= 10; i++ {
		connection.Publish(cache.Context, "pubs", rand.Int())
		time.Sleep(time.Second / 2)
	}

	fmt.Println("Produced!")
}

func consumer(lifecycle chan bool) {
	fmt.Println("Consuming...")
	defer close(lifecycle)

	connection := cache.Client
	topic := connection.Subscribe(cache.Context, "pubs").Channel()

	for {
		select {
		case v := <-topic:
			fmt.Println(v.Payload)
		case <-lifecycle:
			return
		}
	}
}

func main() {
	fmt.Println("Starting...")

	l := make(chan bool)

	go producer(l)
	go consumer(l)

	<-l
	fmt.Println("Out.")
}
