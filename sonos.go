package sonos

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"time"
)

const (
	mx        = 5
	st        = "urn:schemas-upnp-org:device:ZonePlayer:1"
	bcastaddr = "239.255.255.250:1900"
)

type Sonos struct {
	// Context Context
	listenSocket *net.UDPConn
	udpReader    *bufio.Reader
	found        chan *ZonePlayer
}

func NewSonos() (*Sonos, error) {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{IP: []byte{0, 0, 0, 0}, Port: 0, Zone: ""})
	if err != nil {
		return nil, err
	}

	s := Sonos{
		listenSocket: conn,
		udpReader:    bufio.NewReader(conn),
		found:        make(chan *ZonePlayer),
	}

	return &s, nil
}

func (s *Sonos) Close() {
	s.listenSocket.Close()
}

func (s *Sonos) Search() (chan *ZonePlayer, error) {
	go func() {
		for {
			response, err := http.ReadResponse(s.udpReader, nil)
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
				s.found <- zp
			}
		}
	}()

	// MX should be set to use timeout value in integer seconds
	pkt := []byte(fmt.Sprintf("M-SEARCH * HTTP/1.1\r\nHOST: %s\r\nMAN: \"ssdp:discover\"\r\nMX: %d\r\nST: %s\r\n\r\n", bcastaddr, mx, st))
	bcast, err := net.ResolveUDPAddr("udp", bcastaddr)
	if err != nil {
		return nil, err
	}
	_, err = s.listenSocket.WriteTo(pkt, bcast)
	if err != nil {
		return nil, err
	}

	return s.found, nil
}

func FindRoom(room string, timeout time.Duration) (*ZonePlayer, error) {
	son, err := NewSonos()
	if err != nil {
		return nil, err
	}
	defer son.Close()

	found, _ := son.Search()
	to := time.After(timeout)
	for {
		select {
		case <-to:
			return nil, errors.New("timeout")
		case zp := <-found:
			if zp.RoomName() == room {
				return zp, nil
			}
		}
	}
}
