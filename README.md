A little background knolodge on how sonos works will make usage of this API a little eaiser. Sonos used upnp for controll, This is broken down into 4 steps.

1) Simple Service Discovery Discovery Protocol, This step sends a multicast packet to the network on port 1400. Devices thar are upnp enabe will repsond with a HTTP style reposnse that includes a URL that describes the device
2) Simple Device Control Protocol. Using the location in step one, each device can be queried and a list of services supported by the device is returned

These two steps are performed on the `sonos` go library using the Search*() methods

3) Step three involves enumerating examaining each service listed in the scpd document. Each service inslcudes a URL that retunes an xml documnet describing "actions" this service can perform.

4) Perorming a describe action by creating a soap object and sending to to the service control endpoint. 

This library include a sepetate package for each service. 
To use a service, inmport in and call servie.NewService(DeviceUrl)

For example to use the RenderingControl service:

```
import rc "gtihub.com/szatmary/autohome/sonos/renderingcontrol"
control, err := rc.NewRenderingControl("http://device:1400")
```
Methods can be invoked on the retuend object.
Ever methog asspets a struct containing it argumenst, ands responds with a struct.
Som Actions do not take, and arguments, or dont return any arguments. However the methods
Still require an empty struct pass into them.

examples:
```
vol, err := control.GetVolume(&rc.GetVolumeArgs{Channel: "Master"})
fmp.Printf("CurrentVolume %d\n",vol.CurrentVolume)

// In this case SetVolumeResult contains no fields
_, err := control.SetVolume(&rc.SetVolumeArgs{Channel: "Master", DesiredVolume: 25})
```

The service implimentations are automatically generated from the sdcp by the makeservice.go program included.

Different services provide different functionality.
RenderingControl includes action like SetVolume, GetMute, ets. AVTransport allows for starting and stoping playback. Explore the services to see what is possible. 