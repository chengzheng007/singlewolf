package singlewolf

type route struct {
	pattern string //匹配字符串
	handler HandlerFunc
}

func NewRoute(pattern string, handler HandlerFunc) *route {
	return &route{pattern, handler}
}
