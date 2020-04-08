---
title: "meshctl demo istio-multicluster"
weight: 5
---
## meshctl demo istio-multicluster

Install Service Mesh Hub to a cluster, then install/register two Istio installations: one to the management plane cluster, and one to a separate cluster

### Synopsis

Install Service Mesh Hub to a cluster, then install/register two Istio installations: one to the management plane cluster, and one to a separate cluster

```
meshctl demo istio-multicluster [flags]
```

### Options

```
  -h, --help                         help for istio-multicluster
      --remote-cluster-name string   Name under which to register the second cluster
      --remote-context string        In multicluster demos where --use-kind is not specified, use this flag to specify what context to use from your kubeconfig for the remote cluster
```

### Options inherited from parent commands

```
      --cluster-name string     Registers the cluster under this name
      --context string          The main kubeconfig context to use; this will be the context used for management plane installations
      --demo-label string       Optionally label namespaces that are created during the demo with 'solo.io/hub-demo:$DEMO-LABEL' so they can be cleaned up later
      --kube-timeout duration   Specify the default timeout for requests to kubernetes API servers (default 5s)
      --kubeconfig string       Specify the kubeconfig for the current command
  -n, --namespace string        Specify the namespace where Service Mesh Hub resources should be written (default "service-mesh-hub")
      --use-kind                If this is set, use KinD (Kubernetes in Docker) to stand up local clusters; can not set if --context
  -v, --verbose                 Enable verbose mode, which outputs additional execution details that may be helpful for debugging
```

### SEE ALSO

* [meshctl demo](../meshctl_demo)	 - Command line utilities for running/interacting with Service Mesh Hub demos
