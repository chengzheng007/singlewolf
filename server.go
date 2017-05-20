package singlewolf

import (
	"context"
	"net"
	"net/http"
	"time"
)

var httpServers []*http.Server

// start serve http request
func StartServe(handler *router, addrs []string, timeout time.Duration) error {
	if len(addrs) == 0 {
		return nil
	}

	if handler == nil {
		return nil
	}

	httpServers = []*http.Server{}

	for _, addr := range addrs {
		var (
			l   net.Listener
			err error
		)

		if strings.Index(addr, ":") == -1 {
			l, err = net.Listen("unix", addr)
		} else {
			l, err = net.Listen("tcp", addr)
		}

		if err != nil {
			logf("net.Listen() addr(%s) error(%v)\n", addr, err)
			return err
		}

		if err != nil {
			logf("net.Listen(\"tcp\", addr) error(%v)\n", err)
			return err
		}

		server := &http.Server{
			Handler:        handler,
			ReadTimeout:    timeout,
			WriteTimeout:   timeout,
			MaxHeaderBytes: 1 << 20,
		}
		httpServers = append(httpServers, server)

		go func(srv *http.Server, listener net.Listener) {
			logf("start http listen addr: %s", addr)
			if err := srv.Serve(listener); err != nil {
				if err == http.ErrServerClosed {
					return
				}
				logf("srv.Serve(listener) srv(%v) error(%v)", srv, err)
				panic(err)
			}
		}(server, l)
	}

	return nil
}

// close serve
func Close() {
	for _, s := range httpServers {
		ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
		if err := s.Shutdown(ctx); err != nil {
			logf("s.Shutdown(ctx) error(%v)", err)
		}
	}
	httpServers = []*http.Server{}
}
