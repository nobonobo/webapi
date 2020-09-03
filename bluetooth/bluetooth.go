package bluetooth

import (
	"github.com/nobonobo/spago/jsutil"
)

// RequestDevice ...
func RequestDevice(filter map[string]interface{}) (*Device, error) {
	device, err := jsutil.Await(bluetooth.Call("requestDevice", filter))
	if err != nil {
		return nil, err
	}
	return &Device{Value: device}, nil
}
