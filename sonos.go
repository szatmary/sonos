package sonos

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
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
	defer request.Body.Close()
	query := request.URL.Query()
	sn, ok := query["sn"]
	if !ok {
		fmt.Printf("zonePlayer not found")
		response.WriteHeader(404)
		return
	}

	p, ok := s.zonePlayers.Load(sn[0])
	if !ok {
		fmt.Printf("zonePlayer not found")
		response.WriteHeader(404)
		return
	}
	zonePlayer := p.(*ZonePlayer)
	data, err := ioutil.ReadAll(request.Body)
	if err != nil {
		response.WriteHeader(500)
		return
	}

	var events []interface{}
	if request.URL.Path == zonePlayer.AlarmClock.EventEndpoint().Path {
		events = zonePlayer.AlarmClock.ParseEvent(data)
	}
	if request.URL.Path == zonePlayer.AVTransport.EventEndpoint().Path {
		events = zonePlayer.AVTransport.ParseEvent(data)
	}
	if request.URL.Path == zonePlayer.ConnectionManager.EventEndpoint().Path {
		events = zonePlayer.ConnectionManager.ParseEvent(data)
	}
	if request.URL.Path == zonePlayer.ContentDirectory.EventEndpoint().Path {
		events = zonePlayer.ContentDirectory.ParseEvent(data)
	}
	if request.URL.Path == zonePlayer.DeviceProperties.EventEndpoint().Path {
		events = zonePlayer.DeviceProperties.ParseEvent(data)
	}
	if request.URL.Path == zonePlayer.GroupManagement.EventEndpoint().Path {
		events = zonePlayer.GroupManagement.ParseEvent(data)
	}
	if request.URL.Path == zonePlayer.GroupRenderingControl.EventEndpoint().Path {
		events = zonePlayer.GroupRenderingControl.ParseEvent(data)
	}
	if request.URL.Path == zonePlayer.MusicServices.EventEndpoint().Path {
		events = zonePlayer.MusicServices.ParseEvent(data)
	}
	if request.URL.Path == zonePlayer.Queue.EventEndpoint().Path {
		events = zonePlayer.Queue.ParseEvent(data)
	}
	if request.URL.Path == zonePlayer.RenderingControl.EventEndpoint().Path {
		events = zonePlayer.RenderingControl.ParseEvent(data)
	}
	if request.URL.Path == zonePlayer.SystemProperties.EventEndpoint().Path {
		events = zonePlayer.SystemProperties.ParseEvent(data)
	}
	if request.URL.Path == zonePlayer.VirtualLineIn.EventEndpoint().Path {
		events = zonePlayer.VirtualLineIn.ParseEvent(data)
	}
	if request.URL.Path == zonePlayer.ZoneGroupTopology.EventEndpoint().Path {
		events = zonePlayer.ZoneGroupTopology.ParseEvent(data)
	}

	for _, evt := range events {
		zonePlayer.Event(evt)
	}

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
