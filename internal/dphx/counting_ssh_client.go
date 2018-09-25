package dphx

import (
	"log"
	"net"

	"golang.org/x/crypto/ssh"
)

type CountingSSHClient struct {
	*ssh.Client
}

type countedNetConn struct {
	sshClient *CountingSSHClient
	addr      string
	net.Conn
}

func NewCountingSSHClient(sshClient *ssh.Client) *CountingSSHClient {
	return &CountingSSHClient{sshClient}
}

func (c CountingSSHClient) Dial(network, addr string) (net.Conn, error) {
	nc, err := c.Client.Dial(network, addr)
	if err != nil {
		return nil, err
	}

	log.Printf("Dialing to %s via SSH connection %s", addr, sshClient.RemoteAddr())
	return countedNetConn{&c, addr, nc}, nil
}

func (c countedNetConn) Close() error {
	log.Printf("Closing dial to %s via SSH connection %s", c.addr, c.sshClient.RemoteAddr())
	return c.Conn.Close()
}
