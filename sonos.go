package sonos

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"sync"
	"time"
)

const (
	mx        = 5
	st        = "urn:schemas-upnp-org:device:ZonePlayer:1"
	bcastaddr = "239.255.255.250:1900"
)

// type serviceDispatch struct {
// }

type Args struct {
	FoundZonePlayer func(*Sonos, *ZonePlayer)
}

type Sonos struct {
	// Context Context
	udpConn      *net.UDPConn
	httpListener net.Listener
	zonePlayers  sync.Map
	args         Args
}

func NewSonos(args Args) (*Sonos, error) {
	var err error
	s := &Sonos{args: args}

	// create listener for events
	s.httpListener, err = net.Listen("tcp", ":0")
	if err != nil {
		return nil, err
	}

	go func() {
		http.Serve(s.httpListener, s)
	}()
	// Create listener for M-SEARCH
	s.udpConn, err = net.ListenUDP("udp", &net.UDPAddr{IP: []byte{0, 0, 0, 0}, Port: 0, Zone: ""})
	if err != nil {
		return nil, err
	}

	go func() {
		// TODO Dont let this leak, use a contect to shut it down
		udpReader := bufio.NewReader(s.udpConn)
		for {
			response, err := http.ReadResponse(udpReader, nil)
			if err != nil {
				continue
			}
			location, err := url.Parse(response.Header.Get("Location"))
			if err != nil {
				continue
			}
			newZonePlayer, err := NewZonePlayer(location)
			if err != nil {
				continue
			}
			if newZonePlayer.IsCoordinator() {
				zonePlayer, loaded := s.zonePlayers.LoadOrStore(newZonePlayer.SerialNum(), newZonePlayer)
				if !loaded {
					args.FoundZonePlayer(s, zonePlayer.(*ZonePlayer))
				}
				// TODO handle zone players disappearing
			}
		}
	}()

	s.search()
	return s, nil
}

func (s *Sonos) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	// fmt.Printf("%v\n", request)
	// if player, ok := s.zonePlayers.Load(request.URL.Fragment); ok {
	// 	player
	// }

	fmt.Printf("P: %s\nF: %s\n", request.URL.Path, request.URL.Fragment)

	defer request.Body.Close()
	// requestBody, err := ioutil.ReadAll(request.Body)
	// if err != nil {
	// 	response.WriteHeader(500)
	// 	return
	// }

	// fmt.Printf("body: %v\n", string(requestBody))
	response.WriteHeader(200)
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

func FindRoom(room string, timeout time.Duration) (*ZonePlayer, error) {
	c := make(chan *ZonePlayer)
	defer close(c)
	son, err := NewSonos(Args{
		FoundZonePlayer: func(s *Sonos, zp *ZonePlayer) {
			if zp.RoomName() == room {
				c <- zp
			}
		},
	})
	if err != nil {
		return nil, err
	}

	defer son.Close()
	to := time.After(timeout)
	for {
		select {
		case <-to:
			return nil, errors.New("timeout")
		case zp := <-c:
			return zp, nil
		}
	}
	return nil, errors.New("not found")
}
