package lazydialer

import (
	"context"
	"net"
	"reflect"
	"sync"
)

// ContextDialer provides DialContext.
type ContextDialer interface {
	DialContext(ctx context.Context, network, address string) (net.Conn, error)
}

// LazyDialer creates the real dialer on demand.
type LazyDialer struct {
	CreateRealDialer func() (ContextDialer, error)

	real        ContextDialer
	creationErr error

	mu sync.Mutex
}

// DialContext creates a dialer (if it doesn't already exist)
// and dials using it.
func (d *LazyDialer) DialContext(ctx context.Context, network, addr string) (net.Conn, error) {
	real, err := d.getReal()
	if err != nil {
		return nil, err
	}
	if real == nil {
		panic("no real dialer")
	}
	return real.DialContext(ctx, network, addr)
}

func (d *LazyDialer) getReal() (ContextDialer, error) {
	d.mu.Lock()
	defer d.mu.Unlock()

	if !isNil(d.real) {
		return d.real, nil
	}

	if !isNil(d.creationErr) {
		return nil, d.creationErr
	}

	real, err := d.CreateRealDialer()
	if err != nil {
		d.creationErr = err
		return nil, err
	}
	d.real = real
	return real, nil
}

func isNil(i interface{}) bool {
	if i == nil {
		return true
	}

	v := reflect.ValueOf(i)
	return v.Kind() == reflect.Ptr && v.IsNil()
}
