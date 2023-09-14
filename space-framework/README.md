<!--readme-for-space-framework-start-->

<img alt="" width="500px" align="left" src="../KubeStellar-with-Logo.png" />

<br/>
<br/>
<br/>
<br/>

## Space framework 

## Overview
The space framework creates and manages kubernetes api servers by introducing the concepts of spaces and space providers. 
<br>
To utilize the framework you need to [build the space framework from source](#Building), [apply the CRDs](#Applying-the-CRDs), and [start the space manager](#Starting-the-manager). You can then [create providers and spaces](#Creating-providers-and-spaces).

## Building

To build the space framework from scratch issue the following from the space-framework subdirectory:
```shell
$ make codegen
$ make build
```

## Applying the CRDs

To use the space framework you need to create the space and providers CRDs:
```shell
$ kubectl apply -f space-framework/config/crds/space.kubestellar.io_spaceproviderdescs.yaml
$ kubectl apply -f space-framework/config/crds/space.kubestellar.io_spaces.yaml
```

## Starting the manager

To start the manager issue:
```shell
$ ./space-framework/bin/space-manager
```

## Creating providers and spaces

### Creating a Kind provider
A space is created within a provider.  In the following example, we will be using [Kind](https://kind.sigs.k8s.io/), a tool for creating local Kubernetes clusters using Docker containers.  Instructions on how to install and configure Kind can be found [here](https://kind.sigs.k8s.io/docs/user/quick-start/#installation). 
<br>
Here is an example of a provider object which uses Kind as its physical provider:
```shell
$ cat <<EOF | kubectl apply -f -
apiVersion: space.kubestellar.io/v1alpha1
kind: SpaceProviderDesc
metadata:
  name: default
spec:
  ProviderType: "kind"
  SpacePrefixForDiscovery: "ks-"
EOF
```

Once the provider object is created the space manager will create a namespace for it where space objects using the provider will reside. Both the provider object and the namespace can be observed as follows:
```shell
$ kubectl get spaceproviderdescs
NAME      AGE
default   11s
$ kubectl get namespaces
NAME                    STATUS   AGE
spaceprovider-default   Active   12s
```

### Creating a space within a Kind provider
Once you have a provider defined, you can create a space using it. Here is an example of a space object which uses the Kind provider defined above:
```shell
$ cat <<EOF | kubectl apply -f -
apiVersion: space.kubestellar.io/v1alpha1
kind: Space
metadata:
  name: space1
  namespace: "spaceprovider-default"
spec:
  SpaceProviderDescName: "default"
  Managed: true
EOF
```

Once the space object is created, the space manager will detect it and create the corresponding space within the Kind cluster.
```shell
$ kubectl get spaces -A
NAMESPACE               NAME     AGE
spaceprovider-default   space1   10s
$ kind get clusters
space1
```

Note that the space1 created is a managed space, meaning it is created upon the application of a space object and is removed upon the removal of said object.  Alternatively, spaces can be discovered within a provider in which cases they are unmanaged. Spaces that have names that match the provider's SpacePrefixForDiscovery are automatically discovered and an applicable space object is created for them.  

### Importing a space

If you have an existing kubernetes cluster which you wish to use as a space, you can import it into the space framework.  To import a space create a space object that does not have corresponding provider and set the type to imported, as follows. Note that the ClusterConfig contains the cluster configuration as it is found in the kubeconfig file.

```shell
$ cat <<EOF | kubectl apply -f -
apiVersion: space.kubestellar.io/v1alpha1
kind: Space
metadata:
  name: space2
spec:
  Type: imported
status:
  ClusterConfig: |
    apiVersion: v1
    clusters:
    - cluster:
        certificate-authority-data: <your-data-goes-here>
        server: https://127.0.0.1:33625
      name: kind-import1
    contexts:
    - context:
        cluster: kind-import1
        user: kind-import1
      name: kind-import1
    current-context: kind-import1
    kind: Config
    preferences: {}
    users:
    - name: kind-import1
      user:
        client-certificate-data: <your-data-goes-here>
        client-key-data: <your-data-goes-here>
EOF
```

An imported space does not have a provider. An imported space object does not have to reside in its own namespace. Once the object is detected by the space manager, the space will be available for use.

```shell
$ kubectl get space space2 -o yaml | grep Phase
  Phase: Ready
```

<!--readme-for-space-framework-end-->
