<!--kubestellar-syncer-0-deploy-guilder-start-->
Go to inventory management workspace and find the mailbox workspace name.
```shell
espw_space_config="${PWD}/temp-space-config/espw.config"
kubectl-kubestellar-get-config-for-space --space-name espw --provider-name default --sm-core-config $SM_CONFIG --sm-context $SM_CONTEXT --output $espw_space_config

pvname=`kubectl --kubeconfig $espw_space_config get synctargets.edge.kubestellar.io | grep guilder | awk '{print $1}'`
stuid=`kubectl --kubeconfig $espw_space_config get synctargets.edge.kubestellar.io $pvname -o jsonpath="{.metadata.uid}"`
mbws="imw1-mb-$stuid"
echo "mailbox workspace name = $mbws"
```

``` { .bash .no-copy }
Current workspace is "root:imw1".
mailbox workspace name = vosh9816n2xmpdwm-mb-bf1277df-0da9-4a26-b0fc-3318862b1a5e
```

Go to the mailbox workspace and run the following command to obtain yaml manifests to bootstrap KubeStellar-Syncer.
```shell
mbwsname_space_config="${PWD}/temp-space-config/${mbws}.config"
kubectl-kubestellar-get-config-for-space --space-name ${mbws} --provider-name default --sm-core-config $SM_CONFIG --sm-context $SM_CONTEXT --output $mbwsname_space_config

./bin/kubectl-kubestellar-syncer_gen --kubeconfig $mbwsname_space_config guilder --syncer-image quay.io/kubestellar/syncer:v0.2.2 -o guilder-syncer.yaml
```
``` { .bash .no-copy }
Current workspace is "root:vosh9816n2xmpdwm-mb-bf1277df-0da9-4a26-b0fc-3318862b1a5e".
Creating service account "kubestellar-syncer-guilder-wfeig2lv"
Creating cluster role "kubestellar-syncer-guilder-wfeig2lv" to give service account "kubestellar-syncer-guilder-wfeig2lv"

1. write and sync access to the synctarget "kubestellar-syncer-guilder-wfeig2lv"
2. write access to apiresourceimports.

Creating or updating cluster role binding "kubestellar-syncer-guilder-wfeig2lv" to bind service account "kubestellar-syncer-guilder-wfeig2lv" to cluster role "kubestellar-syncer-guilder-wfeig2lv".

Wrote WEC manifest to guilder-syncer.yaml for namespace "kubestellar-syncer-guilder-wfeig2lv". Use

  KUBECONFIG=<workload-execution-cluster-config> kubectl apply -f "guilder-syncer.yaml"

to apply it. Use

  KUBECONFIG=<workload-execution-cluster-config> kubectl get deployment -n "kubestellar-syncer-guilder-wfeig2lv" kubestellar-syncer-guilder-wfeig2lv

to verify the syncer pod is running.
Current workspace is "root:espw".
```

Deploy the generated yaml manifest to the target cluster.
```shell
KUBECONFIG=~/.kube/config kubectl --context kind-guilder apply -f guilder-syncer.yaml
```
``` { .bash .no-copy }
namespace/kubestellar-syncer-guilder-wfeig2lv created
serviceaccount/kubestellar-syncer-guilder-wfeig2lv created
secret/kubestellar-syncer-guilder-wfeig2lv-token created
clusterrole.rbac.authorization.k8s.io/kubestellar-syncer-guilder-wfeig2lv created
clusterrolebinding.rbac.authorization.k8s.io/kubestellar-syncer-guilder-wfeig2lv created
secret/kubestellar-syncer-guilder-wfeig2lv created
deployment.apps/kubestellar-syncer-guilder-wfeig2lv created
```
    
Check that the syncer is running, as follows.
```shell
KUBECONFIG=~/.kube/config kubectl --context kind-guilder get deploy -A
```
``` { .bash .no-copy }
NAMESPACE                             NAME                                  READY   UP-TO-DATE   AVAILABLE   AGE
kubestellar-syncer-guilder-saaywsu5   kubestellar-syncer-guilder-saaywsu5   1/1     1            1           52s
kube-system                           coredns                               2/2     2            2           35m
local-path-storage                    local-path-provisioner                1/1     1            1           35m
```

<!--kubestellar-syncer-0-deploy-guilder-end-->
