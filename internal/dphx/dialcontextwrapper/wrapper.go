package dialcontextwrapper

import (
	"net"

	"golang.org/x/net/context"
)

// NoContextDialer provides an old-style Dial func (with no context).
type NoContextDialer interface {
	Dial(network, addr string) (net.Conn, error)
}

// DialContextWrapper wraps old style no-context Dial to be availabe for
// use as DialContext, where passed context is just ignored.
type DialContextWrapper struct {
	NoContextDialer
}

// DialContext calls NoContextDialer.Dial ignoring the context.
func (d *DialContextWrapper) DialContext(ctx context.Context, network, addr string) (net.Conn, error) {
	// Just ignore context.
	return d.NoContextDialer.Dial(network, addr)
}
