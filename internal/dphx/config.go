package dphx // import "github.com/MOZGIII/dphx/internal/dphx"

// AppConfig loads and saves config.
type AppConfig struct {
	SSH       SSHConfig
	LocalAddr string `envconfig:"SOCKS_ADDR" default:"127.0.0.1:1080"`
}
