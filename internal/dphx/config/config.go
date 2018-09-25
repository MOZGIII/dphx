package config // import "github.com/MOZGIII/dphx/internal/dphx/config"

// App stores application configuration.
type App struct {
	SSH       SSH
	LocalAddr string `envconfig:"SOCKS_ADDR" default:"127.0.0.1:1080"`
}

// SSH stores configuration for the SSH client.
type SSH struct {
	ServerAddr string   `envconfig:"ADDR"`
	Username   string   `envconfig:"USER"`
	Password   string   `envconfig:"PASSWORD"`
	PublicKeys []string `envconfig:"PUBLIC_KEYS"`
	AgentAddr  string   `envconfig:"AGENT"`
}
