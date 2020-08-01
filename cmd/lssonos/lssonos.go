package main

import (
	"fmt"

	sonos "github.com/szatmary/gono6"
)

func main() {
	son, err := sonos.NewSonos()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer son.Close()

	err = son.Search()
	if err != nil {
		return
	}
	// to := time.After(10 * time.Second)
	// for {
	// 	select {
	// 	case <-to:
	// 		return
	// 	case zp := <-found:
	// 		fmt.Printf("%s\t%s\t%s\n", zp.RoomName(), zp.ModelName(), zp.SerialNum())
	// 	}
	// }
}
