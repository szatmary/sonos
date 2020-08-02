package main

import (
	"fmt"
	"time"

	sonos "github.com/szatmary/sonos"
)

func main() {
	son, err := sonos.NewSonos(sonos.Args{
		FoundZonePlayer: func(sonos *sonos.Sonos, player *sonos.ZonePlayer) {
			fmt.Printf("%s\t%s\t%s\n", player.RoomName(), player.ModelName(), player.SerialNum())
			sonos.SubscribeAll(player)
		},
	})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer son.Close()
	time.Sleep(300 * time.Second)
}
