package singlewolf

import (
	"net/http"
)

type Request struct {
	*http.Request
	Params paramsData
}
