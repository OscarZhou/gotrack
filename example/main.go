package main

import track "gotrack"

func main() {
	track := track.Default()
	track.Start()
	defer track.End()
}
