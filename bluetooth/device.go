package bluetooth

import (
	"syscall/js"

	"github.com/nobonobo/spago/jsutil"
)

// Device ...
type Device struct {
	js.Value
	releases []jsutil.Releaser
}

// JSValue for js.Wrapper
func (d *Device) JSValue() js.Value {
	return d.Value
}

// Connect ...
func (d *Device) Connect() (*Server, error) {
	server, err := jsutil.Await(d.Value.Get("gatt").Call("connect"))
	if err != nil {
		return nil, err
	}
	return &Server{Value: server}, nil
}

// OnDisconnect ...
func (d *Device) OnDisconnect(fn func(d *Device)) {
	d.releases = append(d.releases,
		jsutil.Bind(d.Value, "gattserverdisconnected", func(ev js.Value) {
			console.Call("log", "discconnect:", ev)
			fn(d)
		}),
	)
}

// Disconnect ...
func (d *Device) Disconnect() {
	d.Value.Get("gatt").Call("disconnect")
}

// Release ...
func (d *Device) Release() {
	d.Disconnect()
	for _, r := range d.releases {
		r.Release()
	}
}
