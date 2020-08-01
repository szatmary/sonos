package sonnos

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"sync"
)

const (
	mx        = 5
	st        = "urn:schemas-upnp-org:device:ZonePlayer:1"
	bcastaddr = "239.255.255.250:1900"
)

type serviceDispatch struct {
}

type Sonos struct {
	// Context Context
	udpConn      *net.UDPConn
	httpListener net.Listener
	zonePlayers  sync.Map
}

func NewSonos(found func(*ZonePlayer)) (*Sonos, error) {

	// create listener for events
	httpListener, err := net.Listen("tcp", ":0")
	if err != nil {
		return nil, err
	}

	// go func() {
	// 	http.Serve(httpListener, &s)
	// }()
	// Create listener for M-SEARCH
	udpConn, err := net.ListenUDP("udp", &net.UDPAddr{IP: []byte{0, 0, 0, 0}, Port: 0, Zone: ""})
	if err != nil {
		return nil, err
	}

	go func() {
		// TODO Dont let this leak, use a contect to shut it down
		udpReader := bufio.NewReader(udpConn)
		for {
			response, err := http.ReadResponse(udpReader, nil)
			if err != nil {
				continue
			}
			location, err := url.Parse(response.Header.Get("Location"))
			if err != nil {
				continue
			}
			zp, err := NewZonePlayer(location)
			if err != nil {
				continue
			}
			if zp.IsCoordinator() {
				found(zp)
				// fmt.Printf("+++%s\n", GetLocalAddress())
				// err := zp.RenderingControl.Subscribe(zp.HttpClient, GetLocalAddress())
			}
		}
	}()

	s := Sonos{
		zonePlayers:  sync.Map{},
		udpConn:      udpConn,
		httpListener: httpListener,
	}
	s.search()
	return &s, nil
}

func (s *Sonos) HttpPort() int {
	return s.httpListener.Addr().(*net.TCPAddr).Port
}

func (s *Sonos) Close() {
	s.udpConn.Close()
	s.httpListener.Close()
}

func (s *Sonos) search() error {
	// MX should be set to use timeout value in integer seconds
	pkt := []byte(fmt.Sprintf("M-SEARCH * HTTP/1.1\r\nHOST: %s\r\nMAN: \"ssdp:discover\"\r\nMX: %d\r\nST: %s\r\n\r\n", bcastaddr, mx, st))
	bcast, err := net.ResolveUDPAddr("udp", bcastaddr)
	if err != nil {
		return err
	}
	_, err = s.udpConn.WriteTo(pkt, bcast)
	if err != nil {
		return err
	}

	return nil
}

// func FindRoom(room string, timeout time.Duration) (*ZonePlayer, error) {
// 	son, err := NewSonos(func(zp *ZonePlayer) {
// 		// defer son.Close()

// 		// to := time.After(timeout)
// 		// for {
// 		// 	select {
// 		// 	case <-to:
// 		// 		return nil, errors.New("timeout")
// 		// 	case zp := <-son.found:
// 		// 		if zp.RoomName() == room {
// 		// 			return zp, nil
// 		// 		}
// 		// 	}
// 		// }
// 	})

// 	if err != nil {
// 		return nil, err
// 	}
// }
