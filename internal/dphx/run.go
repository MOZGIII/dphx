package dphx // import "github.com/MOZGIII/dphx/internal/dphx"

import (
	"log"

	"github.com/MOZGIII/dphx/internal/dphx/cli"
	"github.com/MOZGIII/dphx/internal/dphx/config"
	"github.com/MOZGIII/dphx/internal/dphx/dialcontextwrapper"
	"github.com/MOZGIII/dphx/internal/dphx/emptyresolver"
	"github.com/MOZGIII/dphx/internal/dphx/lazydialer"
	socks5 "github.com/armon/go-socks5"
)

// Run initializes config and starts SOCKS server.
func Run() {
	if err := boot(); err != nil {
		log.Fatalf("Error: %s", err)
		return
	}
}

func boot() error {
	appConfig, err := config.FromENV()
	if err != nil {
		return err
	}
	cli.PrintEnv(appConfig)

	lazyDialer := lazydialer.LazyDailer{
		CreateRealDailer: func() (lazydialer.ContextDialer, error) {
			sshClient, cerr := createSSHClient(appConfig)
			if cerr != nil {
				return nil, cerr
			}
			dialer := dialcontextwrapper.DialContextWrapper{
				NoContextDialer: sshClient,
			}
			return &dialer, nil
		},
	}

	server, err := socks5.New(&socks5.Config{
		Dial:     lazyDialer.DialContext,
		Resolver: emptyresolver.EmptyResolver{},
	})
	if err != nil {
		return err
	}

	log.Printf("SOCKS5 server is starting at %s", appConfig.LocalAddr)
	return server.ListenAndServe("tcp", appConfig.LocalAddr)
}
