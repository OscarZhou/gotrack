package main

import (
	track "gotrack"
	"log"
	"time"
)

func main() {
	t, err := track.New(track.Config{
		Debug:           true,
		AsynLog:         true,
		AsynLogInterval: 2,
	})
	if err != nil {
		log.Fatal(err)
	}
	c := make(chan int, 2)
	go Loop8(t, c)
	go Loop10(t, c)

	for i := 0; i < 2; i++ {
		// Make sure the main thread won't stop
		// before all threads end
		<-c
	}
}

// Loop8 starts a 8s' loop
func Loop8(t *track.Track, c chan int) {
	t.Start()
	defer func(chan int) {
		c <- 1
	}(c)
	defer t.End()
	for i := 0; i < 8; i++ {
		time.Sleep(time.Duration(1 * time.Second))
	}
}

// Loop10 starts a 10s' loop
func Loop10(t *track.Track, c chan int) {
	t.Start()
	defer func(chan int) {
		c <- 1
	}(c)
	defer t.End()
	for i := 0; i < 10; i++ {
		time.Sleep(time.Duration(1 * time.Second))
	}
}
