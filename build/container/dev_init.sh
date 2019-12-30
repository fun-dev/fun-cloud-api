sleep 5
kind create cluster
export KUBECONFIG="$(kind get kubeconfig-path --name="kind")"
realize start --name='container' --run