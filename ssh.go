package dphx

import (
	"fmt"
	"log"
	"net"

	"golang.org/x/crypto/ssh"
	"golang.org/x/net/context"
)

var sshClient *ssh.Client

func createSSHClient() (*ssh.Client, error) {
	sshConfig, err := appConfig.SSH.ClientConfig()

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

func ensureSSHClient() error {
	if sshClient == nil {
		newClient, err := createSSHClient()
		if err != nil {
			return err
		}
		sshClient = newClient
	}
	return nil
}

// SSHDial does dial via SSH connection to a remote server.
func SSHDial(ctx context.Context, network, addr string) (net.Conn, error) {
	if err := ensureSSHClient(); err != nil {
		return nil, err
	}

	log.Printf("Dialing to %s via SSH connection %s", addr, sshClient.RemoteAddr())

	return sshClient.Dial(network, addr)
}
