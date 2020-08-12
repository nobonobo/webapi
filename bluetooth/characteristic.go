package bluetooth

import (
	"syscall/js"

	"github.com/nobonobo/spago/jsutil"
)

// Characteristic ...
type Characteristic struct {
	js.Value
	releases []jsutil.Releaser
}

// JSValue for js.Wrapper
func (c *Characteristic) JSValue() js.Value {
	return c.Value
}

// Release ...
func (c *Characteristic) Release() {
	for _, r := range c.releases {
		r.Release()
	}
}

// ReadValue ...
func (c *Characteristic) ReadValue() ([]byte, error) {
	data, err := jsutil.Await(c.Value.Call("readValue"))
	if err != nil {
		return nil, err
	}
	return jsutil.JS2Bytes(data), nil
}

// WriteValue ...
func (c *Characteristic) WriteValue(b []byte) error {
	_, err := jsutil.Await(c.Value.Call("writeValue", jsutil.Bytes2JS(b)))
	if err != nil {
		return err
	}
	return nil
}

// StartNotifications ...
func (c *Characteristic) StartNotifications() error {
	_, err := jsutil.Await(c.Value.Call("startNotifications"))
	if err != nil {
		return err
	}
	return nil
}

// OnCharacteristicValueChanged ...
func (c *Characteristic) OnCharacteristicValueChanged(fn func(c *Characteristic)) {
	c.releases = append(c.releases,
		jsutil.Bind(c.Value, "characteristicvaluechanged", func(ev js.Value) {
			console.Call("log", "characteristicvaluechanged:", ev)
			fn(c)
		}),
	)
}
