package gorpc_demo

import (
	"testing"
	"time"
)

func TestService(t *testing.T) {
	 opts := &ServerOptions{
	 	network: "tcp",
	 	address: "127.0.0.1:8000",
	 	timeout: time.Millisecond * 1000,
	 }

	 s := &service{}

	 go func() {
	 	s.Serve(opts)
	 }()

	 s.Close()
}
