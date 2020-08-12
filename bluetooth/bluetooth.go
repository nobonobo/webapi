package bluetooth

import (
	"github.com/nobonobo/spago/jsutil"
)

// Bluetooth ...
type Bluetooth struct {
	devices []*Device
}

// RequestDevice ...
func (bt *Bluetooth) RequestDevice(filter map[string]interface{}) ([]*Device, error) {
	device, err := jsutil.Await(bluetooth.Call("requestDevice", filter))
	if err != nil {
		return nil, err
	}
	res := []*Device{}
	for i := 0; i < device.Length(); i++ {
		res = append(res, &Device{Value: device.Index(i)})
	}
	return res, nil
}

// Release ...
func (bt *Bluetooth) Release() {
	for _, d := range bt.devices {
		d.Release()
	}
}
