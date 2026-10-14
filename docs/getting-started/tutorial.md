# Tutorial: a from-scratch cluster

This is a hands-on walkthrough. By the end you will have a running Kubernetes
cluster on Xen Orchestra, built from the **`almalinux-fromscratch`** template: a
minimal VM image whose nodes install kubelet and kubeadm at bootstrap.

It assumes the provider is already [installed](installation.md) on a management
cluster, the XO credentials secret exists, and you have a VM template ready (see
[Prerequisites](prerequisites.md)).

!!! note "From-scratch or pre-filled?"

    The repository ships two machine images:

    - **`almalinux-fromscratch`** — a minimal image; the node installs
      `kubelet`, `kubeadm` and `containerd` itself at bootstrap. Nothing to
      rebuild when you change the Kubernetes version, but each node installs
      packages on first boot (a bit slower to join).
    - **`almalinux-prefilled`** — kubelet, kubeadm, containerd and the kube-vip /
      Cilium images are baked in; bootstrap is minimal and fast, but the image is
      tied to one Kubernetes version.

    This tutorial uses the from-scratch variant. Swapping to pre-filled is a
    one-line change (see the [kubeadm guide](../guides/providers/kubeadm.md)).

## What you are about to create

A cluster is just a handful of Kubernetes objects. `clusterctl` will render them
from a template, but it is worth knowing what they are:

| Resource | Role |
|---|---|
| `Cluster` | The logical cluster: networking, and references to the control plane and the infrastructure. |
| `XOCluster` | The **infrastructure** cluster: control plane endpoint, load balancer, addons. |
| `XOMachineTemplate` (×2) | The **infrastructure** stencil for control plane and worker VMs: template, pool, network, size. |
| `KubeadmControlPlane` | The control plane provider: how many control plane nodes, and their kubeadm config. |
| `KubeadmConfigTemplate` | The bootstrap config for worker nodes. |
| `MachineDeployment` | The worker group. |
| `MachineHealthCheck` (×2) | Auto-remediation when a node stops being healthy. |

The provider fills in the rest: for every `Machine`, it creates an `XOMachine`,
which becomes a VM.

## 1. Prepare your environment

Export the values for your infrastructure. `clusterctl` reads them from the
environment when it renders the template:

```bash
export CP_HOST=192.168.1.240        # the control plane VIP (must be free)
export CP_PORT=6443
export CP_LB=kube-vip               # or "" for no load balancer
export CP_SUBNET=16                 # CIDR bits of the VIP subnet
export VM_NAME_PREFIX=my-cluster
export XO_TEMPLATE_UUID=<your-vm-template-uuid>   # the from-scratch template
export XO_POOL_UUID=<your-pool-uuid>
export XO_NETWORK_UUID=<your-network-uuid>
# export XO_CNI=none                # default is cilium
```

`CP_HOST` is the address the API server will be reachable at. With `kube-vip` it
is a floating IP that the control plane nodes share — it must be free and routed.

## 2. Render the manifests

Generate the cluster objects into a file so you can read them before applying:

```bash
clusterctl generate cluster my-cluster \
  --from templates/kubeadm/base/clusterctl/almalinux-fromscratch.yaml \
  --kubernetes-version v1.36.1 \
  --control-plane-machine-count 1 \
  --worker-machine-count 1 > my-cluster.yaml
```

Open `my-cluster.yaml`. You will recognise the resources from the table above:
the `XOCluster` carries your `CP_HOST` and `controlPlaneLB`, the two
`XOMachineTemplate`s carry the template, pool and network, and the
`KubeadmControlPlane` / `MachineDeployment` wire everything together.

!!! tip "One Kubernetes version, everywhere"

    `--kubernetes-version` must be the version the nodes will actually run.
    The from-scratch template pins exactly that version at bootstrap; if KCP and
    the node disagree, the machine is rolled forever. See the
    [kubeadm guide](../guides/providers/kubeadm.md) for the details.

## 3. Apply

```bash
kubectl apply -f my-cluster.yaml
```

## 4. Watch it come up

The provider now creates the VMs, hands them their bootstrap configuration, and
waits for them to join:

```bash
kubectl get xocluster my-cluster -w      # the cluster-level infrastructure
kubectl get xomachines -w                # one per VM
kubectl get machines -w                  # the generic view
```

In Xen Orchestra you will see VMs named `my-cluster-…` appear, cloned from your
template, configured, and started. When a machine is ready its
`status.providerID` is set to `xenorchestra://<pool-uuid>/<vm-uuid>` and
`status.addresses` holds its IP.

!!! note "First boot is slow with from-scratch"

    Each node installs kubelet, kubeadm and containerd before it can join. Give
    the first node a few minutes; watch the controller logs if you are curious:

    ```bash
    kubectl -n capi-system logs -l app.kubernetes.io/name=vates-capi -f
    ```

## 5. Reach the cluster

Once the control plane is up, fetch its kubeconfig:

```bash
clusterctl get kubeconfig my-cluster > my-cluster.kubeconfig
kubectl --kubeconfig my-cluster.kubeconfig get nodes
```

If the nodes stay `NotReady`, the workload cluster has no CNI yet. The shipped
template installs **Cilium** through a `ClusterResourceSet`; if you set
`XO_CNI=none`, install a CNI of your choice.

## 6. Clean up

```bash
kubectl delete -f my-cluster.yaml
```

Deleting the objects cascades: the provider shuts the VMs down and removes them
from Xen Orchestra. Watch it happen, or follow the logs:

```bash
kubectl -n capi-system logs -l app.kubernetes.io/name=vates-capi -f
```

## Where next

- Prefer a managed topology you can evolve declaratively? The
  [ClusterClass flow](../guides/providers/kubeadm.md) uses an overlay instead of
  a flat template.
- Want an immutable, cloud-init-free OS? Try the
  [Talos flow](../guides/providers/talos/index.md).
- Tune the VMs (CPU, RAM, disk, networks) and the cluster (addons, SSH keys):
  see [Reference → Custom resources](../reference/crds/xocluster.md).
