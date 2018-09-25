package emptyresolver

import (
	"net"

	"golang.org/x/net/context"
)

// EmptyResolver provides a resolver that does nothing.
type EmptyResolver struct{}

// Resolve simply returns the passed context, nil and no error.
func (EmptyResolver) Resolve(ctx context.Context, name string) (context.Context, net.IP, error) {
	return ctx, nil, nil
}
