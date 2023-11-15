#!/usr/bin/env bash

# Copyright 2023 The KubeStellar Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -e

KUBESTELLAR_SERVICE="kubestellar"


function echoerr() {
   echo "ERROR: $1" >&2
}

function wait_kcp_ready() {
    echo "Waiting for kcp to be ready... this may take a while."
    (
        until [ "$(kubectl --kubeconfig $host_kubeconfig logs $(kubectl --kubeconfig $host_kubeconfig get pod --selector=app=kubestellar -o jsonpath='{.items[0].metadata.name}') -c kcp | grep '***READY***')" != "" ]; do
           sleep 10
        done
    )
}

function get_kcp_kubeconfig() {
    wait_kcp_ready
    kcp_kubeconfig_dir=/home/kubestellar/.kcp
    kcp_kubeconfig=${kcp_kubeconfig_dir}/admin.kubeconfig
    echo "Copying the admin.kubeconfig from kubestellar seret..."
    mkdir -p $kcp_kubeconfig_dir
    (
        kubectl --kubeconfig $host_kubeconfig get secrets kubestellar -o 'go-template={{index .data "admin.kubeconfig"}}' | base64 --decode > $kcp_kubeconfig
    )
}

function create_kcp_provider() {
    get_kcp_kubeconfig     # kcp is created in a seperate container 
    echo "Delete the kcp provider secret if it already exists, and then create it."
    if ! kubectl --kubeconfig $SPACE_MANAGER_KUBECONFIG delete secret -n ${NAMESPACE} kcpsec ; then
        echo "Nothing to delete."
    fi
    kubectl --kubeconfig $host_kubeconfig get secrets kubestellar -o 'go-template={{index .data "admin.kubeconfig"}}' | base64 --decode > kcpsecret
    kubectl --kubeconfig $SPACE_MANAGER_KUBECONFIG create secret generic -n ${NAMESPACE} kcpsec --from-file=kubeconfig="kcpsecret"    
    rm kcpsecret

    echo "Delete the kcp provider object if it already exists, and then create it."
    if ! kubectl --kubeconfig $SPACE_MANAGER_KUBECONFIG delete spaceproviderdesc $PROVIDER_NAME ; then
        echo "Nothing to delete."
    fi
    kubectl --kubeconfig $SPACE_MANAGER_KUBECONFIG create -f - <<EOF
apiVersion: space.kubestellar.io/v1alpha1
kind: SpaceProviderDesc
metadata:
  name: $PROVIDER_NAME
spec:
  ProviderType: "kcp"
  SpacePrefixForDiscovery: "ks-"
  secretRef:
    namespace: ${NAMESPACE}
    name: kcpsec
EOF
}
 
function create_kubeflex_provider() {
    echo "Waiting for the kubeflex provider to be ready... this may take a while."
    (
        until [ "$(kubectl --kubeconfig $host_kubeconfig get pods -A | grep kubeflex-controller-manager | grep Running)" != "" ]; do
            sleep 10
        done
    )

    echo "Delete the kubeflex provider object if it already exists, and then create it."
    if ! kubectl --kubeconfig $SPACE_MANAGER_KUBECONFIG delete spaceproviderdesc $PROVIDER_NAME ; then
        echo "Nothing to delete."
    fi
    kubectl --kubeconfig $SPACE_MANAGER_KUBECONFIG create -f - <<EOF
apiVersion: space.kubestellar.io/v1alpha1
kind: SpaceProviderDesc
metadata:
  name: $PROVIDER_NAME
spec:
  ProviderType: "kubeflex"
  SpacePrefixForDiscovery: "ks-"
  secretRef:
    namespace: ${NAMESPACE}
    name: corecluster
EOF
    echo "Waiting for default spaceprovider to reach the Ready phase."
    until [ "$(kubectl --kubeconfig ${SPACE_MANAGER_KUBECONFIG} get spaceproviderdesc $PROVIDER_NAME -o yaml | grep Ready )" != "" ]; do
        sleep 1
    done 
}


function set_provider_adapters() {
    echo "Waiting for space manager to be ready... this may take a while."
    (
        until [ "$(kubectl --kubeconfig $host_kubeconfig logs $(kubectl --kubeconfig $host_kubeconfig get pod --selector=app=kubestellar -o jsonpath='{.items[0].metadata.name}') -c space-manager | grep '***READY***')" != "" ]; do
            sleep 10
        done
    )
    if [ "$SPACE_PROVIDER_TYPE" == "kcp" ]; then
        create_kcp_provider
    elif [ "$SPACE_PROVIDER_TYPE" == "kubeflex" ]; then
        create_kubeflex_provider
    else
        echo "No valid default space provider."
    fi
}

function wait-kubestellar-ready() {
    echo "Waiting for KubeStellar to be ready... this may take a while."
    (
        until [ "$(kubectl --kubeconfig $host_kubeconfig logs $(kubectl --kubeconfig $host_kubeconfig get pod --selector=app=kubestellar -o jsonpath='{.items[0].metadata.name}') -c init | grep '***READY***')" != "" ]; do
            sleep 10
        done
    )
    echo "Success!"
}

function guess_kcp_dns() {
    if [ -z "$EXTERNAL_HOSTNAME" ]; then
        # Try to guess the route
        if kubectl --kubeconfig $host_kubeconfig get route kubestellar-route &> /dev/null; then
            EXTERNAL_HOSTNAME=$(kubectl --kubeconfig $host_kubeconfig get route kubestellar-route -o yaml -o jsonpath={.spec.host} 2> /dev/null)
        fi
    fi
    if [ -z "$EXTERNAL_HOSTNAME" ]; then
        # Try to guess the ingress
        if kubectl --kubeconfig $host_kubeconfig get ingress kubestellar-ingress &> /dev/null; then
            EXTERNAL_HOSTNAME=$(kubectl --kubeconfig $host_kubeconfig get ingress kubestellar-ingress -o yaml -o jsonpath={.spec.rules[0].host} 2> /dev/null)
        fi
    fi
    echo "${EXTERNAL_HOSTNAME}"
}

function run_kcp() {
    echo "--< Starting kcp >--"

    echo Attempting to delete kubestellar secret...
    (
        KUBECONFIG=$host_kubeconfig
        if ! kubectl delete secret kubestellar ; then
            echo "Nothing to delete."
        fi
    )
    echo "EXTERNAL_HOSTNAME=${EXTERNAL_HOSTNAME}"

    # Check EXTERNAL_HOSTNAME
    if [ -z "$EXTERNAL_HOSTNAME" ]; then
        echo "Trying to guess the DNS from route/ingress...."
        export EXTERNAL_HOSTNAME=$(guess_kcp_dns)
    fi
    echo "EXTERNAL_HOSTNAME=${EXTERNAL_HOSTNAME}"

    # Create the certificates
    if [ -n "$EXTERNAL_HOSTNAME" ]; then
        echo "Creating the TLS certificates"
        # mkdir -p .kcp
        cd .kcp
        eval pieces_external=($(kubestellar-ensure-kcp-server-creds ${EXTERNAL_HOSTNAME}))
        eval pieces_cluster=($(kubestellar-ensure-kcp-server-creds ${KUBESTELLAR_SERVICE})) #############
        cd ..
    fi

    # Running kcp
    if [ -n "$EXTERNAL_HOSTNAME" ]; then
         # required to fix the restart
        echo "Removing existing apiserver keys... "
        if ! rm /home/kubestellar/.kcp/apiserver.* &> /dev/null ; then
            echo "Nothing to remove... must be the first time."
        else
            echo "Existing keys removed... the container mast have restarted."
        fi
        echo -n "Running kcp with TLS keys... "
        kcp start --tls-sni-cert-key ${pieces_external[1]},${pieces_external[2]} --tls-sni-cert-key ${pieces_cluster[1]},${pieces_cluster[2]} & # &> kcp.log &
    else
        echo -n "Running kcp without TLS keys... "
        kcp start &
    fi
    echo Started.

    # Waiting to be ready
    echo "Waiting for ${KUBECONFIG}..."
    while [ ! -f "${KUBECONFIG}" ]; do
        sleep 5;
    done
    echo 'Waiting for "root:compute" workspace...'
    until [ "$(kubectl ws root:compute 2> /dev/null)" != "" ]; do
        sleep 5;
    done
    echo '"root:compute" workspace is ready'.
    echo "kcp version: $(kubectl version --short 2> /dev/null | grep kcp | sed 's/.*kcp-//')"
    kubectl ws root

    # Generate the external.kubeconfig and cluster.kubeconfig
    if [ -n "$EXTERNAL_HOSTNAME" ] && [ ! -d "${PWD}/.kcp-${EXTERNAL_HOSTNAME}" ]; then
        echo Creating external.kubeconfig...
        switch-domain .kcp/admin.kubeconfig .kcp/external.kubeconfig root ${EXTERNAL_HOSTNAME} ${EXTERNAL_PORT} ${pieces_external[0]}
        switch-domain .kcp/admin.kubeconfig .kcp/cluster.kubeconfig root ${KUBESTELLAR_SERVICE} 6443 ${pieces_cluster[0]}
    fi

    # Ensure kubeconfig secret
    echo Creating the kubestellar secret...
    (
        KUBECONFIG=$host_kubeconfig
        if [ -n "${EXTERNAL_HOSTNAME}" ]; then
            kubectl create secret generic kubestellar --from-file="${PWD}/.kcp/admin.kubeconfig" --from-file="${PWD}/.kcp/cluster.kubeconfig" --from-file="${PWD}/.kcp/external.kubeconfig"
        else
            kubectl create secret generic kubestellar --from-file="${PWD}/.kcp/admin.kubeconfig"
        fi
    )

    touch ready
    while true ; do
        echo "***READY***"
        sleep 600
    done
}

function run_init() {
    echo "--< Starting init >--"
    set_provider_adapters
    kubestellar init --ensure-imw $ENSURE_IMW --ensure-wmw $ENSURE_WMW 
    touch ready
    echo "***READY***"
    sleep infinity
}

function run_mailbox_controller() {
    echo "--< Starting mailbox-controller >--"
    wait-kubestellar-ready
    # TODO: remove the kcp dependency
    get_kcp_kubeconfig
    KUBECONFIG=$kcp_kubeconfig
    kubectl ws root:espw
    if ! mailbox-controller -v=${VERBOSITY} ; then
        echoerr "unable to start mailbox-controller!"
        exit 1
    fi
}

function run_where_resolver() {
    echo "--< Starting where-resolver >--"
    wait-kubestellar-ready
    # TODO: remove the kcp dependency
    get_kcp_kubeconfig
    KUBECONFIG=$kcp_kubeconfig
    kubectl ws root:espw
    if ! kubestellar-where-resolver -v ${VERBOSITY} ; then
        echoerr "unable to start kubestellar-where-resolver!"
        exit 1
    fi
}

function run_placement_translator() {
    echo "--< Starting placement-translator >--"
    wait-kubestellar-ready
    # TODO: remove the kcp dependency
    get_kcp_kubeconfig
    KUBECONFIG=$kcp_kubeconfig
    kubectl ws root:espw
    if ! placement-translator --allclusters-context  "system:admin" -v=${VERBOSITY} ; then
        echoerr "unable to start mailbox-controller!"
        exit 1
    fi
}

# get_host_kubeconfig: 
# The hosting cluster is by default the kubestellar and space core cluster. 
function get_host_kubeconfig() {
    kubectl --kubeconfig $host_kubeconfig config set-cluster space-mgt --server="https://${KUBERNETES_SERVICE_HOST}:${KUBERNETES_SERVICE_PORT}" --certificate-authority=/var/run/secrets/kubernetes.io/serviceaccount/ca.crt
    kubectl --kubeconfig $host_kubeconfig config set-credentials space-mgt --token="$(cat /var/run/secrets/kubernetes.io/serviceaccount/token)"
    kubectl --kubeconfig $host_kubeconfig config set-context space-mgt --cluster=space-mgt --user=space-mgt
    kubectl --kubeconfig $host_kubeconfig config use-context space-mgt
}

echo "--< Starting KubeStellar container >--"

export KUBECONFIG_DIR="space-config"
mkdir -p $KUBECONFIG_DIR
host_kubeconfig="${KUBECONFIG_DIR}/config"
get_host_kubeconfig

echo "Environment variables:"
if [ $# -ne 0 ] ; then
    ACTION="$1"
else
    ACTION="sleep"
fi
echo "ACTION=${ACTION}"
if [ "$VERBOSITY" == "" ]; then
    VERBOSITY="2"
fi
if [ "$ESPW_NAME" == "" ]; then
    ESPW_NAME="espw"
fi
if [ "$NAMESPACE" == "" ]; then
    NAMESPACE="default"
fi
if [ "$PROVIDER_NAME" == "" ]; then
    PROVIDER_NAME="default"
    export PROVIDER_NAME
fi
if [ "$PROVIDER_NAMESPACE" == "" ]; then
    PROVIDER_NAMESPACE=spaceprovider-${PROVIDER_NAME}
    export PROVIDER_NAMESPACE
fi
if [ "$SPACE_MANAGER_KUBECONFIG" == "" ]; then
    # if the space_manager_kubeconfig is not set, then we assume the host
    # cluster is the space manager cluster.
    SPACE_MANAGER_KUBECONFIG=$host_kubeconfig
    export SPACE_MANAGER_KUBECONFIG
fi
if [ "$IN_CLUSTER" == "" ]; then
    IN_CLUSTER=true
    export IN_CLUSTER
fi
echo "ESPW_NAME=${ESPW_NAME}"
echo "VERBOSITY=${VERBOSITY}"
echo "ENSURE_IMW=${ENSURE_IMW}"
echo "ENSURE_WMW=${ENSURE_WMW}"
echo "NAMESPACE=${NAMESPACE}"
echo "SPACE_PROVIDER_TYPE=${SPACE_PROVIDER_TYPE}"
echo "KUBECONFIG_DIR=${KUBECONFIG_DIR}"
echo "SPACE_MANAGER_KUBECONFIG=${SPACE_MANAGER_KUBECONFIG}"
echo "PROVIDER_NAME=${PROVIDER_NAME}"
echo "PROVIDER_NAMESPACE=${PROVIDER_NAMESPACE}"
echo "IN_CLUSTER=${IN_CLUSTER}"

case "${ACTION}" in

(kcp)
    run_kcp;;
(init)
    run_init;;
(mailbox-controller)
    run_mailbox_controller;;
(where-resolver)
    run_where_resolver;;
(placement-translator)
    run_placement_translator;;
(sleep)
    echo "Nothing to do... sleeping forever."
    sleep infinity;;
(*)
    echoerr "unknown action '$1'!"
    exit 1;;
esac
