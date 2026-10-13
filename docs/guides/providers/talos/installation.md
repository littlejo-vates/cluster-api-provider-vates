# Talos installation

The Talos providers add `TalosControlPlane` and `TalosConfig` resources, and
they need RBAC access to vates-capi's `XOMachineTemplate` resources, which the
default bindings do not grant.

!!! note "Draft — not written yet"

    Planned contents:

    - `clusterctl init --bootstrap talos --control-plane talos`
    - Why the default vates RBAC only binds the kubeadm control plane
    - `kubectl apply -k config/rbac/talos` and what it grants
    - Installing vates-capi alongside the Talos providers
