# XOCluster

`XOCluster` is the *infrastructure cluster* for a workload cluster. It holds the
settings that apply to the whole cluster rather than to an individual machine,
and it publishes the control plane endpoint once it is known.

One `XOCluster` exists per workload cluster. Cluster API creates it for you when
you create a `Cluster` whose `spec.infrastructureRef` points at it.

## `spec`

| Field | Type | Default | Description |
|---|---|---|---|
| `controlPlaneEndpoint` | `APIEndpoint` | — | The reachable API server endpoint. |
| `controlPlaneLB` | `string` | — | Control plane load balancer. Only `kube-vip` is supported. |
| `injectSSHKeys` | `bool` | `false` | Merge the XO user's SSH keys into cloud-init. |
| `identityRef` | `LocalObjectReference` | — | A `Secret` with the XO credentials for this cluster. |
| `addons` | `AddonsSpec` | CCM + CSI on, no CNI | What the provider installs into the workload cluster. |

### `controlPlaneEndpoint`

| Field | Type | Description |
|---|---|---|
| `host` | `string` | DNS name or IP of the API server. With kube-vip, this is the floating VIP. |
| `port` | `int32` | API server port (`6443` by convention). |
| `subnet` | `int32` | CIDR subnet mask for the load balancer. Auto-detected if unset. |

### `addons`

Installed via `ClusterResourceSet`. Each addon is optional.

| Field | Type | Default | Description |
|---|---|---|---|
| `ccm` | `*bool` | `true` | Xen Orchestra cloud-controller-manager (node addresses, service load balancers). |
| `csi` | `*bool` | `true` | Xen Orchestra CSI driver (persistent volumes). |
| `cni` | `*string` | `none` | Container network interface. Supported: `none`, `cilium`. |
| `nodeOutOfService` | `NodeOutOfServiceSpec` | enabled | CCM out-of-service controller (see below). |

`nodeOutOfService` applies the `node.kubernetes.io/out-of-service` taint to a
node whose VM is deleted, halted, paused, or suspended and whose kubelet reports
`NotReady`. That lets the Kubernetes controller manager force-detach the node's
volumes instead of waiting for its six-minute timer, so persistent volumes can
be reused promptly.

| Field | Type | Default | Description |
|---|---|---|---|
| `enabled` | `*bool` | `true` | Toggle the controller. |
| `syncPeriod` | `*Duration` | CCM default | How often nodes are reconciled. |
| `gracePeriod` | `*Duration` | CCM default | How long a powered-off VM must stay down before the taint is applied. A VM deleted from XO is tainted immediately. |

## `status`

| Field | Type | Description |
|---|---|---|
| `ready` | `bool` | The cluster infrastructure is ready. |
| `controlPlaneEndpoint` | `APIEndpoint` | The endpoint once discovered. |
| `initialization.provisioned` | `bool` | Part of the CAPI contract; drives provisioning order. |
| `failureReason` / `failureMessage` | `*string` | Set on terminal failure. |
| `conditions` | `[]metav1.Condition` | Standard conditions. |

## Example

```yaml
apiVersion: vates.infrastructure.cluster.x-k8s.io/v1beta2
kind: XOCluster
metadata:
  name: my-cluster
spec:
  controlPlaneEndpoint:
    host: 192.168.1.240
    port: 6443
    subnet: 16
  controlPlaneLB: kube-vip
  injectSSHKeys: true
  addons:
    ccm: true
    csi: true
    cni: cilium
```

!!! note "Draft — examples to extend"

    This page documents the fields today. More worked examples (minimal cluster,
    multi-tenant with `identityRef`, disabling addons) will be added.
