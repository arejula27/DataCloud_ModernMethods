#!/bin/bash

# Check if kubectl is installed, if not, install it.
if ! command -v kubectl &> /dev/null
then
    echo "kubectl not found, installing..."
    curl -LO "https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl"
    chmod +x kubectl
    sudo mv kubectl /usr/local/bin/
fi

# Check if helm is installed, if not, install it.
if ! command -v helm &> /dev/null
then
    echo "helm not found, installing..."
    curl https://raw.githubusercontent.com/helm/helm/master/scripts/get-helm-3 | bash
fi

# Add the Argo Helm repository to Helm.
helm repo add argo https://argoproj.github.io/argo-helm

# Install Argo using Helm.
helm install argo argo/argo
