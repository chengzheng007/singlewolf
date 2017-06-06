package singlewolf

import (
	"errors"
	"net/http"
)

/**
孤狼: 专门用于解析简单的json body体请求
只允许post方式提交
*/

var (
	// ErrURIInvalid is patten 错误
	ErrURIInvalid = errors.New("uri pattern error")
	// ErrURIRepeat is repeat 错误
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
