# ClusterClass

Instead of a flat set of manifests, you can describe a cluster as a **managed
topology**: a `Cluster` with a `topology` block, built from a `ClusterClass`.
Cluster API then derives every machine and template for you, and can roll them
out in a controlled way when you change the class.

With vates-capi, this relies on the `XOClusterTemplate` and `XOMachineTemplate`
templates (and their bootstrap counterparts `KubeadmConfigTemplate`).

!!! note "Draft — not written yet"

    Planned contents:

    - When to use a ClusterClass versus a flat template
    - The shipped classes (`almalinux-prefilled`, `almalinux-fromscratch`)
    - The `base/` + `overlays/` layout
    - Editing a managed topology and watching the rollout
    - The ClusterTopology feature gate prerequisite
