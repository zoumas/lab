# Install Go

Download a binary release suitable for your system.

Latest Stable version: `go1.21.3.linux-amd64.tar.gz`

Go installation steps

1. Remove any previous Go installation

by deleting the `/usr/local/go` folder (if it exists), then extract the archive you 
just downloaded into `/usr/local`, creating a fresh Go tree in `/usr/local/go`:

```md
$ sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.21.3.linux-amd64.tar.gz
```

2. Add `/usr/local/go/bin` to the PATH environment variable.

You can do this by adding the following line to your `$HOME/.profile` or `/etc/profile` (for a system-wide installation)
(or `$HOME/.bashrc` - the shell's configuration file):

```md
export PATH=$PATH:/usr/local/go/bin
```

Refresh your current shell session to see the changes with either of these commands:
```md
source $HOME/.bashrc
exec bash -l
```

3. Verify that you've installed Go by typing the following command:
```md
$ go version
```

