package cli // import "github.com/MOZGIII/dphx/internal/dphx/cli"

import (
	"fmt"

	"github.com/MOZGIII/dphx/internal/dphx/config"
)

// PrintEnv prints current environment.
func PrintEnv(appConfig *config.App) {
	fmt.Printf("ENV:\n")
	printAppConfig(appConfig)
	fmt.Printf("\n")
}

func printAppConfig(cfg *config.App) {
	print("SSH Server Address", cfg.SSH.ServerAddr)
	print("SSH Username", cfg.SSH.Username)

	if cfg.SSH.Password == "" {
		print("SSH Password", "<empty>")
	} else {
		print("SSH Password", "[hidden]")
	}

	printArray("SSH Public Keys", cfg.SSH.PublicKeys)
	print("SSH Agent Address", cfg.SSH.AgentAddr)
	print("SOCKS5 Server Address", cfg.LocalAddr)
}

func print(key, value string) {
	fmt.Printf("  %-24s %s\n", key+":", value)
}

func printArray(key string, values []string) {
	fmt.Printf("  %-24s\n", key+":")

	if len(values) == 0 {
		fmt.Printf("    <empty>\n")
	} else {
		for _, value := range values {
			fmt.Printf("    -  %s\n", value)
		}
	}
}
