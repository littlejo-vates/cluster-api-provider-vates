# Prerequisites

Before installing vates-capi, make sure the following are in place. The list is
short, but two items are easy to forget: **cloud-init support** in your VM
template, and the **credentials secret**.

## On your workstation

- **`kubectl`** configured against your management cluster.
- **[`clusterctl`](https://cluster-api.sigs.k8s.io/clusterctl/overview.html)** —
  the Cluster API CLI. You will use it to install the providers and to generate
  cluster manifests.

## A management cluster

A Kubernetes cluster where Cluster API will run. Any of these works:

- [kind](https://kind.sigs.k8s.io/) — the quickest way to get one locally.
- [k3s](https://k3s.io/) — lightweight, good on a small VM.
- Any existing Kubernetes cluster you can apply manifests to.

The management cluster can be small: it only runs the CAPI controllers, not your
workloads.

## Xen Orchestra access

You need an account able to create VMs, plus the identifiers of the resources
the provider will use:

| What | Where to find it | Used for |
|---|---|---|
| **XO URL** | Your Xen Orchestra web UI address | Connecting the provider |
| **XO API token** | XO → user profile → *Authentication tokens* | Connecting the provider |
| **VM template UUID** | XO → the template's *Advanced* tab | Cloning nodes |
| **Pool UUID** | XO → the pool's page | Choosing where VMs are created |
| **Network UUID** | XO → the network's page | Attaching VM interfaces |

### About the VM template

The template is the most common source of trouble, so it deserves its own note.
It must be:

- **Bootable**, with the OS already installed. It is *not* a bare ISO.
- **Cloud-init enabled.** This example uses a standard distribution: the provider
  delivers the bootstrap configuration through cloud-init, so a template that
  ignores it boots and then sits there doing nothing.
- **Running Xen guest tools.** The provider reads the VM's addresses through the
  guest agent; without it, the machine never reports an address and the node
  never joins.
- **Not tied to a specific pool**, so it can be cloned wherever `poolID` points.

!!! tip "Not using cloud-init?"

    The Talos flow uses a `nocloud` template and no cloud-init at all. If that is
    your case, read the [Talos guide](../guides/providers/talos/index.md) instead — its
    prerequisites are different.

## The credentials secret

The provider reads Xen Orchestra credentials from a Kubernetes `Secret` in its
own namespace. It looks for it **at startup and exits if it is missing**, so
create it before installing the provider:

```bash
kubectl create namespace capi-system
kubectl create secret generic xo-credentials -n capi-system \
  --from-literal=url="https://<your-xoa>" \
  --from-literal=token="<your-xo-token>" \
  --from-literal=insecure="true"
```

The secret must contain the keys `url` and `token`, and may contain `insecure`
(set it to `"true"` only if your XO instance uses a self-signed certificate).

!!! warning "Namespace matters"

    The default provider namespace is `capi-system`. If you install the provider
    elsewhere, create the secret in that same namespace, or point the controller
    at another secret with `--xoa-secret-name=<name>`.

    Individual clusters and machines can override these credentials with their
    own `identityRef`, which is handy for multi-tenant setups.

Next: [install the provider](installation.md).
