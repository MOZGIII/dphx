package dphx

import (
	"fmt"
	"log"

	"github.com/MOZGIII/dphx/internal/dphx/authmethods"
	"github.com/MOZGIII/dphx/internal/dphx/config"
	"golang.org/x/crypto/ssh"
)

func createSSHClient(appConfig *config.App) (*ssh.Client, error) {
	auths, err := authmethods.FromConfig(&appConfig.SSH)
	if err != nil {
		return nil, err
	}

	sshConfig := &ssh.ClientConfig{
		User:            appConfig.SSH.Username,
		Auth:            auths,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	if err != nil {
		return nil, err
	}

	log.Printf("Connecting to SSH at %s", appConfig.SSH.ServerAddr)

	client, err := ssh.Dial("tcp", appConfig.SSH.ServerAddr, sshConfig)
	if err != nil {
		return nil, fmt.Errorf("Failed to dial to SSH server: %s", err.Error())
	}

	return client, nil
}
