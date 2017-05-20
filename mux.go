package singlewolf

import (
	"net/http"
)

type mux struct {
	RTR *router
}

func NewMux() *mux {
	return &mux{
		RTR: nil,
	}
}

func (m *mux) MakeHandler(rt *router) http.Handler {

	m.RTR = rt
	f := m.RTR.getHandlerFunc()
	return wrapHandleFunc(f)
}
