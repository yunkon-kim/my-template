# My-template

Describe a project later


## Execution and development environment

- Operating system (OS): 
    - Ubuntu 22.04
- Languages: 
    - Go: 1.21
    - Python: 3.8.10
- Container runtime:
    - Docker: 20.10.12


## How to run my-template

### Source code based installation and execution

#### Configure build environment

1. Install dependencies

```bash
# Ensure that your system is up to date
sudo apt update -y

# Ensure that you have installed the dependencies, 
# such as `ca-certificates`, `curl`, and `gnupg` packages.
sudo apt install make gcc git
```
2. Install Go

Note - **Install the latest stable version of Go** for my-template contribution/development since backward compatibility is supported.
For example, install Go 1.21.4, which is stable version on 2023-11-30, even though `go.mod` says `go 1.19`. (In the opposite case, you will encounter a build error.)

Example - Install Go 1.21.4, see [Go all releases](https://golang.org/dl/) and [Download and install](https://go.dev/doc/install)

```bash
# Set Go version
GO_VERSION=1.21.4

# Get Go archive
wget https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz

# Remove any previous Go installation and
# Extract the archive into /usr/local/
sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz

# Append /usr/local/go/bin to .bashrc
echo 'export PATH=$PATH:/usr/local/go/bin:$HOME/go/bin' >> ~/.bashrc
echo 'export GOPATH=$HOME/go' >> ~/.bashrc

# Apply the .bashrc changes
source ~/.bashrc

# Verify the installation
echo $GOPATH
go version

```

#### Download source code

Clone my-template repository

```bash
git clone https://github.com/yunkon-kim/my-template.git ${HOME}/my-template
```

#### Build my-template

Build my-template source code

```bash
cd ${HOME}/my-template
make
```

(Optional) Update Swagger API document
```bash
cd ${HOME}/my-template
make swag
```

If you got an error because of missing `swag`, install `swag`:
```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

#### Run my-template binary

Set enviroment variable (See [setup.env](https://github.com/yunkon-kim/my-template/blob/main/conf/setup.env)) 

```bash
source ./conf/setup.env
```

Run my-template server

```bash
cd ${HOME}/my-template
make run
```

#### Health-check my-template

Check if my-template is running

```bash
curl http://localhost:8056/my-template/health

# Output if it's running successfully
# {"message":"my-template API server is running"}
```