package dphx // import "github.com/MOZGIII/dphx/internal/dphx"

import (
	"io/ioutil"
	"log"
	"net"
	"os"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
)

// SSHConfig stores configuration for SSH client
type SSHConfig struct {
	ServerAddr string   `envconfig:"ADDR"`
	Username   string   `envconfig:"USER"`
	Password   string   `envconfig:"PASSWORD"`
	PublicKeys []string `envconfig:"PUBLIC_KEYS"`
	AgentAddr  string   `envconfig:"AGENT"`
}

// SetSSHConfigDefaults adds default values that can only
// be set at runtime.
func SetSSHConfigDefaults(sshConfig *SSHConfig) {
	if sshConfig.AgentAddr == "" {
		sshConfig.AgentAddr = os.Getenv("SSH_AUTH_SOCK")
	}
}

// ClientConfig creates ssh.ClientConfig object from current configuration.
func (s SSHConfig) ClientConfig() (*ssh.ClientConfig, error) {
	auths, err := appConfig.SSH.AuthMethods()

	if err != nil {
		return nil, err
	}

	sshConfig := &ssh.ClientConfig{
		User: appConfig.SSH.Username,
		Auth: auths,
	}

	return sshConfig, nil
}

// AuthMethods creates ssh.AuthMethod objects from current configuration.
func (s SSHConfig) AuthMethods() ([]ssh.AuthMethod, error) {
	auths := []ssh.AuthMethod{}
	var err error

	if auths, err = s.addAuthsForKeys(auths); err != nil {
		return nil, err
	}
	if auths, err = s.addAuthsForPassword(auths); err != nil {
		return nil, err
	}

	return auths, nil
}

func (s SSHConfig) addAuthsForPassword(auths []ssh.AuthMethod) ([]ssh.AuthMethod, error) {
	auths = append(auths, ssh.Password(s.Password))
	return auths, nil
}

func (s SSHConfig) addAuthsForKeys(auths []ssh.AuthMethod) ([]ssh.AuthMethod, error) {
	signers := []ssh.Signer{}
	var err error

	if signers, err = s.addSignersForAgent(signers); err != nil {
		return nil, err
	}
	if signers, err = s.addSignersForKeys(signers); err != nil {
		return nil, err
	}

	auths = append(auths, ssh.PublicKeys(signers...))
	return auths, nil
}

func (s SSHConfig) addSignersForAgent(signers []ssh.Signer) ([]ssh.Signer, error) {
	if s.AgentAddr == "" {
		log.Println("WARNING: SSH agent connection not configured, not using agent auth")
		return signers, nil
	}

	sock, err := net.Dial("unix", s.AgentAddr)
	if err != nil {
		return nil, err
	}

	agent := agent.NewClient(sock)

	agentSigners, err := agent.Signers()
	if err != nil {
		return nil, err
	}

	signers = append(signers, agentSigners...)
	return signers, nil
}

func (s SSHConfig) addSignersForKeys(signers []ssh.Signer) ([]ssh.Signer, error) {
	for _, path := range s.PublicKeys {
		pemBytes, err := ioutil.ReadFile(path)
		if err != nil {
			return nil, err
		}

		signer, err := ssh.ParsePrivateKey(pemBytes)
		if err != nil {
			return nil, err
		}

		signers = append(signers, signer)
	}

	return signers, nil
}
