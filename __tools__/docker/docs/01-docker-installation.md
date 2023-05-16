# Docker Install

[GO TO ORIGIN SITE](https://docs.docker.com/)

## # Ubuntu Bionic 18.04 (LTS)

### - Uninstall old version

```sh
sudo apt-get remove docker docker-engine docker.io containerd runc
```

<br><br><br>

### - Set up the repository

1. Update the `apt` package index and install packages to allow `apt` to use a repository over HTTPS:

```sh
sudo apt-get update
```

```sh
sudo apt-get install \
    ca-certificates \
    curl \
    gnupg \
    lsb-release
```

2. Add Docker’s official GPG key:

```sh
sudo mkdir -p /etc/apt/keyrings
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg
```

3. Use the following command to set up the repository:

```sh
echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu \
  $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
```

<br><br><br>

### - Install Docker Engine

1. Update the `apt` package index, and install the latest version of Docker Engine, containerd, and Docker Compose, or go to the next step to install a specific version:

```sh
sudo apt-get update
```
```sh
sudo apt-get install docker-ce docker-ce-cli containerd.io docker-compose-plugin
```

2. Verify that Docker Engine is installed correctly by running the `hello-world` image.

```sh
sudo docker run hello-world
```

<br><br><br>

### - (Optional) Uninstall Docker Engine

1. Uninstall the Docker Engine, CLI, Containerd, and Docker Compose packages:

```sh
sudo apt-get purge docker-ce docker-ce-cli containerd.io docker-compose-plugin
```

2. Images, containers, volumes, or customized configuration files on your host are not automatically removed. To delete all images, containers, and volumes:

```sh
sudo rm -rf /var/lib/docker
```
```sh
sudo rm -rf /var/lib/containerd
```

<hr><br><br><br><br>

## # CentOS 7

### - Uninstall old versions

```sh
sudo yum remove docker \
                  docker-client \
                  docker-client-latest \
                  docker-common \
                  docker-latest \
                  docker-latest-logrotate \
                  docker-logrotate \
                  docker-engine
```

The contents of `/var/lib/docker/`, including images, containers, volumes, and networks, are preserved. The Docker Engine package is now called `docker-ce`

<br><br><br>

### - Set up the repository

Install the `yum-utils` package (which provides the `yum-config-manager` utility) and set up the repository.

```sh
sudo yum install -y yum-utils
```
```sh
sudo yum-config-manager \
    --add-repo \
    https://download.docker.com/linux/centos/docker-ce.repo
```

<br><br><br>

### - Install Docker Engine

1. Install the latest version of Docker Engine, containerd, and Docker Compose or go to the next step to install a specific version:

```sh
sudo yum install docker-ce docker-ce-cli containerd.io docker-compose-plugin
```

If prompted to accept the GPG key, verify that the fingerprint matches `060A 61C5 1B55 8A7F 742B 77AA C52F EB6B 621E 9F35`, and if so, accept it.

2. Start Docker

```sh
sudo systemctl start docker
```

3. Verify that Docker Engine is installed correctly by running the `hello-world` image.

```sh
sudo docker run hello-world
```

if `WSL2` then `sudo service docker start`

<br><br><br>

### - Uninstall Docker Engine

1. Uninstall the Docker Engine, CLI, Containerd, and Docker Compose packages:

```sh
sudo yum remove docker-ce docker-ce-cli containerd.io docker-compose-plugin
```

2. Images, containers, volumes, or customized configuration files on your host are not automatically removed. To delete all images, containers, and volumes:

```sh
sudo rm -rf /var/lib/docker
sudo rm -rf /var/lib/containerd
```

<br><br><br>
<hr><hr><hr>
<br><br><br>

# Docker Compose Install

## # Installation

1. To download and install the Compose CLI plugin, run:

```sh
sudo curl -L "https://github.com/docker/compose/releases/download/v2.10.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
```

2. Apply executable permissions to the binary:

```sh
sudo chmod +x /usr/local/bin/docker-compose
```
```sh
sudo ln -s /usr/local/bin/docker-compose /usr/bin/docker-compose
```

3. Test the installation.

```sh
docker compose version
```

<br><br><br>

## # Uninstallation

```sh
sudo rm /usr/local/bin/docker-compose
```

<hr><br><br><br><br><br>

## # Manage Docker as a non-root user
To run Compose as a non-root user

 - Warning !!!

```
The docker group grants privileges equivalent to the root user. For details on how this impacts security in your system
```

 1. Create the docker group.

```sh
sudo groupadd docker
```

 2. Add your user to the docker group.

```sh
sudo usermod -aG docker $USER
```

 - [Run the Docker daemon as a non-root user (Rootless mode)](https://docs.docker.com/engine/security/rootless/)