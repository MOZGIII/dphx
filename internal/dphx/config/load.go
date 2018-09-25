package config

import (
	"os"

	"github.com/kelseyhightower/envconfig"
)

// FromENV initializes the the config from env.
func FromENV() (*App, error) {
	appConfig := App{}
	envconfig.MustProcess("DPHX", &appConfig)
	envconfig.MustProcess("DPHX_SSH", &appConfig.SSH)
	if appConfig.SSH.AgentAddr == "" {
		appConfig.SSH.AgentAddr = os.Getenv("SSH_AUTH_SOCK")
	}
	return &appConfig, nil
}
