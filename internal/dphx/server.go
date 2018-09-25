package dphx // import "github.com/MOZGIII/dphx/internal/dphx"

import (
	"github.com/armon/go-socks5"
)

var appConfig AppConfig

// ListenAndServe starts the SOCKS server.
func ListenAndServe(network, addr string) error {
	server, err := createSocks5Server()
	if err != nil {
		return err
	}

	if err := server.ListenAndServe(network, addr); err != nil {
		return err
	}

	return nil
}

func createSocks5Server() (*socks5.Server, error) {
	conf := &socks5.Config{
		Dial:     SSHDial,
		Resolver: EmptyResolver{},
	}

	server, err := socks5.New(conf)
	if err != nil {
		return nil, err
	}
	return server, nil
}
