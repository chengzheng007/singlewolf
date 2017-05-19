package singlewolf


import (
	"fmt"
	"net"
	"net/http"
	"testing"
	"time"
)

var (
	addr        = "127.0.0.1:9088"
	httpTimeout = time.Second * 15
)

func TestA(t *testing.T) {

	l, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Printf("net.Listen(\"tcp\", addr) error(%v)\n", err)
		return
	}

	httpServeMux, err := InitServerMux()
	if err != nil {
		fmt.Printf("InitServerMux() error(%v)\n", err)
		return
	}

	server := &http.Server{Handler: httpServeMux, ReadTimeout: httpTimeout, WriteTimeout: httpTimeout}

	fmt.Println("start ...")

	if err := server.Serve(l); err != nil {
		fmt.Printf("server.Serve(l) error(%v)\n", err)
		return
	}
}

func InitServerMux() (http.Handler, error) {
	mux := NewMux()

	router, err := MakeRouter(
		NewRoute("/hello", hello),
	)

	if err != nil {
		return nil, err
	}

	mux.SetMux(router)

	return mux.MakeHandler(), nil
}

func hello(w *Wrapper, res Result) {
	res["ret"] = "123"

	res["params"] = w.Request.Params.GetAll()

	res["a"] = w.Request.Params.GetString("a")

	// if err := w.ResponseWriter.WriteJson(res); err != nil {
	// 	fmt.Printf("w.ResponseWriter.WriteJson(%v) error(%v)\n", err)
	// }

}
