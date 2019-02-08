package main

import (
	track "gotrack"
	"time"
)

func main() {
	t := track.Default()
	c := make(chan int)
	go Loop5(t, c)

	// Make sure the main thread won't stop
	// before all threads end
	<-c
}

// Loop5 starts a 5s loop
func Loop5(t *track.Track, c chan int) {
	t.Start()
	defer t.End()
	for i := 0; i < 5; i++ {
		time.Sleep(time.Duration(1 * time.Second))
	}
	c <- 1
}
