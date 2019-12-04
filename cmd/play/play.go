package main

import (
	"fmt"
	"os"
	"time"

	"github.com/szatmary/sonos"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Usage: %s [room name] [media url]\n", os.Args[0])
		return
	}

	zp, err := sonos.FindRoom(os.Args[1], 5*time.Second)
	if err != nil {
		fmt.Printf("FindRoom Error: %v\n", err)
		return
	}

	if err = zp.SetAVTransportURI(os.Args[2]); err != nil {
		fmt.Printf("SetAVTransportURI Error: %v\n", err)
		return
	}

	if err = zp.Play(); err != nil {
		fmt.Printf("Play Error: %v\n", err)
		return
	}
}
