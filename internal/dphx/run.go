package dphx // import "github.com/MOZGIII/dphx/internal/dphx"

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

// Run initializes config and starts SOCKS server.
func Run() {
	SetEnv()
	PrintEnv()

	log.Printf("SOCKS5 server is starting at %s", appConfig.LocalAddr)
	if err := ListenAndServe("tcp", appConfig.LocalAddr); err != nil {
		log.Fatalln(err.Error())
	}
}

// SetEnv loads current env config.
func SetEnv() {
	envconfig.MustProcess("DPHX", &appConfig)
	envconfig.MustProcess("DPHX_SSH", &appConfig.SSH)
	SetSSHConfigDefaults(&appConfig.SSH)
}
