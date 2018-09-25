package authmethods

import (
	"io/ioutil"
	"log"
	"net"

	"github.com/MOZGIII/dphx/internal/dphx/config"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
)

// FromConfig provides auth methods from config.
func FromConfig(cfg *config.SSH) ([]ssh.AuthMethod, error) {
	signer, err := signerAuthMethod(cfg)
	if err != nil {
		return nil, err
	}

	return []ssh.AuthMethod{
		signer,
		ssh.Password(cfg.Password),
	}, nil
}

func signerAuthMethod(cfg *config.SSH) (ssh.AuthMethod, error) {
	signers := []ssh.Signer{}

	{
		if addr := cfg.AgentAddr; addr != "" {
			newSigners, err := signersForAgent(addr)
			if err != nil {
				return nil, err

			}
			signers = append(signers, newSigners...)
		} else {
			log.Println("WARNING: SSH agent connection not configured, not using agent auth")
		}
	}

	{
		newSigners, err := signersForPublicKeyFiles(cfg.PublicKeys)
		if err != nil {
			return nil, err
		}
		signers = append(signers, newSigners...)
	}

	return ssh.PublicKeys(signers...), nil
}

func signersForAgent(agentAddr string) ([]ssh.Signer, error) {
	sock, err := net.Dial("unix", agentAddr)
	if err != nil {
		return nil, err
	}

	agent := agent.NewClient(sock)
	return agent.Signers()
}

func signersForPublicKeyFiles(publicKeys []string) ([]ssh.Signer, error) {
	signers := []ssh.Signer{}
	for _, path := range publicKeys {
		pemBytes, err := ioutil.ReadFile(path) // nolint: gas, gosec
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
