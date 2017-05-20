package singlewolf

import (
	"encoding/json"
	"net/http"
)

type ResponseWriter interface {

	// Identical to the http.ResponseWriter interface
	Header() http.Header

	// Use EncodeJson to generate the payload, write the headers with http.StatusOK if
	// they are not already written, then write the payload.
	// The Content-Type header is set to "application/json", unless already specified.
	WriteJson(v interface{}) error

	// Encode the data structure to JSON, mainly used to wrap ResponseWriter in
	// middlewares.
	EncodeJson(v interface{}) ([]byte, error)

	// Similar to the http.ResponseWriter interface, with additional JSON related
	// headers set.
	WriteHeader(int)
}

type responseWriter struct {
	http.ResponseWriter
	wroteHeader bool
}

func (w *responseWriter) Header() http.Header {
	return w.ResponseWriter.Header()
}

func (w *responseWriter) WriteJson(v interface{}) error {
	buf, err := w.EncodeJson(v)
	if err != nil {
		return err
	}
	_, err = w.Write(buf)
	return err
}

func (w *responseWriter) EncodeJson(v interface{}) ([]byte, error) {
	buf, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func (w *responseWriter) WriteHeader(code int) {
	if w.Header().Get("Content-Type") == "" {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
	}
	w.ResponseWriter.WriteHeader(code)
	w.wroteHeader = true
}

func (w *responseWriter) Write(v []byte) (int, error) {
	if !w.wroteHeader {
		w.WriteHeader(http.StatusOK)
	}
	return w.ResponseWriter.Write(v)
}


func notFoundHandle(w http.ResponseWriter) {
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	return
}