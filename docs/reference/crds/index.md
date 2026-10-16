# Custom resources

vates-capi adds four custom resources to the management cluster, all in the API
group **`vates.infrastructure.cluster.x-k8s.io/v1beta2`**:

| Kind | Plural | Scope | Purpose |
|---|---|---|---|
| [`XOCluster`](xocluster.md) | `xoclusters` | Namespaced | Cluster-wide infrastructure settings. |
| [`XOClusterTemplate`](xoclustertemplate.md) | `xoclustertemplates` | Namespaced | Stencil for `XOCluster`, used by ClusterClass. |
| [`XOMachine`](xomachine.md) | `xomachines` | Namespaced | One VM (one node). |
| [`XOMachineTemplate`](xomachinetemplate.md) | `xomachinetemplates` | Namespaced | Stencil that machines are stamped from. |

They are part of the Cluster API "infrastructure" contract: `XOCluster` plays
the role of an *infrastructure cluster*, and `XOMachine` the role of an
*infrastructure machine*. You rarely create them by hand. Instead you write a
`Cluster` and `Machine`/`MachineDeployment`/`KubeadmControlPlane` that point at
the templates, and the provider fills in the rest.

!!! info "Generated from the source"

    These pages mirror the Go types in `api/v1beta2/`. When a field changes, the
    generated CRDs (`config/crd/bases/`) change with `make manifests`, and this
    documentation should follow. If you spot a drift, the source of truth is
    always the `*_types.go` files.

!!! tip "Two-field pattern (ID or name)"

    Wherever a resource can be identified either by its Xen Orchestra UUID or by
    its human name, the two are separate fields — for example `poolID` and
    `poolName`. The controller prefers the `ID` field, which needs no API call.
    Set one or the other, not an ambiguous mix.
