package main

import (
	"fmt"
	"time"

	"github.com/szatmary/sonos"
)

func main() {
	son, err := sonos.NewSonos()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer son.Close()

	found, _ := son.Search()
	to := time.After(10 * time.Second)
	for {
		select {
		case <-to:
			return
		case zp := <-found:
			fmt.Printf("%s\t%s\t%s", zp.RoomName(), zp.ModelName(), zp.SerialNum())
		}
	}
}
