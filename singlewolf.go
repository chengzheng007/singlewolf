package singlewolf

import (
	"errors"
	"net/http"
)

// singlewolf: a json body request processor, it only allow post request
// this package can be used in some scenes like: client serialize all data
// to json to request web API, for example the mobile app Android and IOS.
// They all have similar request way: a raw string data, so we can process
// them in an uniform way

var (
	// ErrURIInvalid imply uri is invalid, uri should be not empty with a leading "/"
	ErrURIInvalid = errors.New("uri pattern error")
	// ErrURIRepeat imply there has been registered a same pattern
	ErrURIRepeat = errors.New("uri is repeat")
)

// HandlerFunc function type to process your business
type HandlerFunc func(*Wrapper, Result)

// Request is 返回信息
type Request struct {
	*http.Request
	Params paramsData
}

// Result write back to front end
type Result map[string]interface{}

// Wrapper wrap request and response,
type Wrapper struct {
	Request
	ResponseWriter
}
