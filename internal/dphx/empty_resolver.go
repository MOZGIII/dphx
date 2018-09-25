package dphx // import "github.com/MOZGIII/dphx/internal/dphx"

import (
	"net"

	"golang.org/x/net/context"
)

type EmptyResolver struct{}

func (EmptyResolver) Resolve(ctx context.Context, name string) (context.Context, net.IP, error) {
	return ctx, nil, nil
}
