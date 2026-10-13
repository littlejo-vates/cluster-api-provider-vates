# XOMachine

`XOMachine` represents **one virtual machine**, and therefore one node of a
workload cluster. Its spec describes the VM to create; its status reports what
was created and whether it is ready.

The provider creates an `XOMachine` for you whenever a `Machine`'s
`spec.infrastructureRef` points at one. You normally only touch `XOMachine`
through `XOMachineTemplate`.

## `spec`

| Field | Type | Description |
|---|---|---|
| `providerID` | `*string` | Set by the controller after creation, as `xenorchestra://<pool-uuid>/<vm-uuid>`. |
| `poolID` / `poolName` | `string` | Pool to create the VM in. Falls back to the template's pool. |
| `templateID` / `templateName` | `string` | VM template to clone. |
| `namePrefix` | `string` | **Required.** Prefix for the VM name in Xen Orchestra. |
| `networkConfig` | `NetworkConfig` | Interfaces, IP filtering, and guest network config. |
| `resourceSet` | `ResourceSet` | vCPUs, memory, disk. |
| `failureDomain` | `*string` | Failure domain this machine belongs to. |
| `bootstrapProvider` | `string` | **Observability only.** Applied as a `bootstrap:<name>` tag on the VM. Does not change behavior. |
| `bootstrap` | `BootstrapBehavior` | What to *do* with the bootstrap payload. See [Bootstrap payload](../bootstrap.md). |
| `bootstrapData` | `string` | Inline bootstrap data. If empty, read from the owning `Machine`'s bootstrap secret. |
| `identityRef` | `LocalObjectReference` | Per-machine XO credentials. Falls back to the cluster's, then the controller's. |

### `networkConfig`

| Field | Type | Description |
|---|---|---|
| `networks` | `[]Network` | Interfaces to attach. Each has `networkID` (UUID) or `name`. |
| `allowedIPRange` | `string` | CIDR of IPs allowed to reach the VM, applied to all its interfaces. Unrestricted if empty. |
| `guestConfig` | `string` | Cloud-init netplan (v2) passed to the guest, for a static IP or custom setup before first boot. |

### `resourceSet`

| Field | Type | Description |
|---|---|---|
| `cpus` | `*int32` | Number of vCPUs (minimum 1). |
| `memory` | `string` | Memory in Kubernetes resource format, e.g. `8Gi`. |
| `diskSize` | `string` | Disk size in Kubernetes resource format, e.g. `30Gi`. |

### `bootstrap`

| Field | Type | Description |
|---|---|---|
| `format` | `string` | `cloudinit` or `passthrough` (default). See [Bootstrap payload](../bootstrap.md). |

Leaving `bootstrap` unset is the same as `format: passthrough`.

## `status`

| Field | Type | Description |
|---|---|---|
| `ready` | `bool` | The VM is up and the node has joined. |
| `providerID` | `*string` | The VM's Xen Orchestra provider ID. |
| `addresses` | `[]corev1.NodeAddress` | The VM's addresses, read through the guest agent. |
| `initialization.provisioned` | `bool` | Part of the CAPI contract. |
| `failureReason` / `failureMessage` | `*string` | Set on terminal failure. |
| `conditions` | `[]metav1.Condition` | Standard conditions. |

## Example

```yaml
apiVersion: vates.infrastructure.cluster.x-k8s.io/v1beta2
kind: XOMachine
metadata:
  name: my-cluster-control-plane-abcde
spec:
  namePrefix: my-cluster
  templateID: 3f2b1c0a-...        # your VM template UUID
  poolID: 9d8c7b6a-...
  resourceSet:
    cpus: 4
    memory: 8Gi
    diskSize: 50Gi
  networkConfig:
    networks:
      - networkID: 1a2b3c4d-...
  bootstrap:
    format: cloudinit
```

!!! note "Draft — examples to extend"

    More examples (multiple interfaces, static IP via `guestConfig`, per-machine
    `identityRef`) will be added.
