package bluetooth

import (
	"syscall/js"

	"github.com/nobonobo/spago/jsutil"
)

// Service ...
type Service struct {
	js.Value
}

// JSValue for js.Wrapper
func (s *Service) JSValue() js.Value {
	return s.Value
}

// GetCharacteristic ...
func (s *Service) GetCharacteristic(uuid string) (*Characteristic, error) {
	char, err := jsutil.Await(s.Value.Call("getCharacteristic", uuid))
	if err != nil {
		return nil, err
	}
	return &Characteristic{Value: char}, nil
}
