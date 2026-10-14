# XOClusterTemplate

A stencil for `XOCluster`, used by [ClusterClass](../../guides/clusterclass.md)
managed topologies. It wraps an `XOCluster` spec under
`spec.template.spec`, exactly like the built-in infrastructure templates do for
other providers.

!!! note "Draft — not written yet"

    Planned contents:

    - The `spec.template.spec` shape, and why it exists
    - How a `ClusterClass` references it
    - Editing the template and the resulting rollout
