<!--required-packages-start-->
!!! tip "Required Packages:"
    === "Mac"
        ``` title="docker - https://docs.docker.com/engine/install/"
        brew install docker
        open -a Docker
        ```
        ``` title="kind - https://kind.sigs.k8s.io/docs/user/quick-start/"
        brew install kind
        ```
        ``` title="kubectl - https://kubernetes.io/docs/tasks/tools/ (version range expected: 1.23-1.25)"
        brew install kubectl
        ```
        ``` title="jq - https://stedolan.github.io/jq/download/"
        brew install jq
        ```
        [GO v1.19](https://gist.github.com/jniltinho/8758e15a9ef80a189fce) - You will need GO to compile and run kcp and the KubeStellar scheduler.  Currently kcp requires go version 1.19.
    === "Ubuntu"
        ``` title="docker - https://docs.docker.com/engine/install/"
        sudo apt-get install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
        ```
        ``` title="kind - https://kind.sigs.k8s.io/docs/user/quick-start/"
        curl -Lo ./kind https://kind.sigs.k8s.io/dl/v0.14.0/kind-linux-amd64; chmod +x ./kind; mv ./kind /usr/local/bin
        ```
        ``` title="kubectl - https://kubernetes.io/docs/tasks/tools/ (version range expected: 1.23-1.25)"
        curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl; chmod +x kubectl; mv ./kubectl /usr/local/bin/kubectl
        ```
        ``` title="jq - https://stedolan.github.io/jq/download/"
        sudo apt-get install jq
        ```
        [GO v1.19](https://gist.github.com/jniltinho/8758e15a9ef80a189fce) - You will need GO to compile and run kcp and the KubeStellar scheduler.  Currently kcp requires go version 1.19.
    === "Centos7"
        ``` title="docker - https://docs.docker.com/engine/install/"
        yum -y install epel-release ; yum -y install docker ; systemctl enable --now docker ;systemctl status docker
        ```
        ``` title="kind - https://kind.sigs.k8s.io/docs/user/quick-start/"
        curl -Lo ./kind https://kind.sigs.k8s.io/dl/v0.14.0/kind-linux-amd64; chmod +x ./kind; mv ./kind /usr/local/bin
        ```
        ``` title="kubectl - https://kubernetes.io/docs/tasks/tools/ (version range expected: 1.23-1.25)"
        curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl; chmod +x kubectl; mv ./kubectl /usr/local/bin/kubectl
        ```
        ``` title="jq - https://stedolan.github.io/jq/download/"
        yum -y install jq
        ```
        [GO v1.19](https://gist.github.com/jniltinho/8758e15a9ef80a189fce) - You will need GO to compile and run kcp and the KubeStellar scheduler.  Currently kcp requires go version 1.19.
    === "Windows"
        ``` title="docker - https://docs.docker.com/engine/install/"
        choco install docker -y
        ```
        ``` title="kind - https://kind.sigs.k8s.io/docs/user/quick-start/"
        curl.exe -Lo kind-windows-amd64.exe https://kind.sigs.k8s.io/dl/v0.14.0/kind-windows-amd64
        ```
        ``` title="kubectl - https://kubernetes.io/docs/tasks/tools/install-kubectl-windows/ (version range expected: 1.23-1.25)"
        curl.exe -LO "https://dl.k8s.io/release/v1.27.2/bin/windows/amd64/kubectl.exe"
        ```
        ``` title="jq - https://stedolan.github.io/jq/download/"
        choco install jq -y
        choco install curl -y
        ```
        [GO v1.19](https://gist.github.com/jniltinho/8758e15a9ef80a189fce) - You will need GO to compile and run kcp and the KubeStellar scheduler.  Currently kcp requires go version 1.19.
<!--required-packages-end-->
<!-- 
## 
  - [docker](https://docs.docker.com/engine/install/)
  - [kind](https://kind.sigs.k8s.io/)
  - [kubectl](https://kubernetes.io/docs/tasks/tools/) (version range expected: 1.23-1.25)
  - [jq](https://stedolan.github.io/jq/download/) -->