package singlewolf

import (
	"net/http"
)

type mux struct {
	RTR MuxI
}

func NewMux() *mux {
	return &mux{
		RTR: nil,
	}
}

func (m *mux) SetMux(router MuxI) {
	m.RTR = router
}

func (m *mux) MakeHandler() http.Handler {
	var f HandlerFunc
	if m.RTR != nil {
		f = m.RTR.MuxFunc()
	} else {
		f = func(*Wrapper, Result) {}
	}

	return wrapHandleFunc(f)
}
