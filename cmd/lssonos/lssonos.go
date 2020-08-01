package main

import (
	"fmt"
	"time"

	sonos "github.com/szatmary/sonos"
)

func main() {
	son, err := sonos.NewSonos(func(zp *sonos.ZonePlayer) {
		fmt.Printf("%s\t%s\t%s\n", zp.RoomName(), zp.ModelName(), zp.SerialNum())
	})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer son.Close()
	time.Sleep(10 * time.Second)
}
