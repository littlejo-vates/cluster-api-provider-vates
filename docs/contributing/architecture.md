# Architecture

A map of the codebase: where to look for what, and how a reconciliation actually
flows through the packages.

!!! note "Draft — not written yet"

    Planned contents:

    - Repository layout (`api/`, `cmd/`, `internal/`, `config/`, `templates/`)
    - The controllers: `XOCluster` and `XOMachine` reconciliation
    - The `internal/kubevip` injection, and the bootstrap payload handling
    - The Xen Orchestra client and how it is constructed
    - The addon installers (CCM, CSI, CNI) via `ClusterResourceSet`
