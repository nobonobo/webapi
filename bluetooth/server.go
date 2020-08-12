package bluetooth

import (
	"syscall/js"

	"github.com/nobonobo/spago/jsutil"
)

// Server ...
type Server struct {
	js.Value
	releases []jsutil.Releaser
}

// JSValue for js.Wrapper
func (s *Server) JSValue() js.Value {
	return s.Value
}

// Release ...
func (s *Server) Release() {
	for _, r := range s.releases {
		r.Release()
	}
}

// GetPrimaryService ...
func (s *Server) GetPrimaryService(uuid string) (*Service, error) {
	service, err := jsutil.Await(s.Value.Call("getPrimaryService", uuid))
	if err != nil {
		return nil, err
	}
	return &Service{Value: service}, nil
}
