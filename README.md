# dphx

[![Build Status](https://travis-ci.org/MOZGIII/dphx.svg?branch=master)](https://travis-ci.org/MOZGIII/dphx)

An SSH client that provides a tunnel to the remote network via a local SOCKS 5 server.

Designed to be managed as an unprivileged system service and work all the time at the background. Will only establish connection to SSH host if needed (i.e. when SOCKS request arrives) and disconnect automatically after some inactivity period. At the same time, SOCKS server is always listening for requests, and will manage SSH connection appropriately.

Remote DNS resolution is supported.

First SSH connection happens on first SOCKS request (requiring it to exist, but everything is proxified right now, so that would be any request).

## Installation

### GitHub Releases

The simplest way to install this software is to download a binary for your system from the [GitHub Releases](https://github.com/MOZGIII/dphx/releases).

### Ubuntu PPA

[`dphx` PPA](https://launchpad.net/~mozgiii/+archive/ubuntu/dphx) is available for Ubuntu users.

```shell
sudo add-apt-repository ppa:mozgiii/dphx
sudo apt-get update
```

### Go tools

To install via **Go** toolchain, use the following command:

```
go get github.com/MOZGIII/dphx/cmd/dphx
```

## Config

Config is loaded from the process environment.

A sample settings file along with options description and example values is listed below.

```bash
# SSH server address.
export DPHX_SSH_ADDR='example.net'

# SSH username.
export DPHX_SSH_USER='user'

# SSH password (optional).
export DPHX_SSH_PASSWORD='password'

# SSH public key files list (optional).
export DPHX_SSH_PUBLIC_KEYS='/home/me/.ssh/id_rsa,/home/me/.ssh/other_id_rsa'

# SSH agent address, usually a unix socket.
# Defaults to $SSH_AUTH_SOCK value.
export DPHX_SSH_AGENT='/run/user/1337/keyring-blah/ssh'

# SOCKS5 server bind address (optional).
# Defaults to "127.0.0.1:1080".
export DPHX_SOCKS_ADDR='0.0.0.0:1080'
```
