# XOMachineTemplate

A stencil for `XOMachine`, referenced by `KubeadmControlPlane`,
`MachineDeployment`, and other machine-producing objects. Instead of writing one
`XOMachine` per node, you write one `XOMachineTemplate` and let the control
plane or deployment stamp out machines from it.

It wraps an `XOMachine` spec under `spec.template.spec`.

!!! note "Draft — not written yet"

    Planned contents:

    - The `spec.template.spec` shape
    - Referencing it from a `KubeadmControlPlane` / `MachineDeployment`
    - Rolling changes: edit the template to reshape future machines
    - The RBAC access the Talos control plane provider needs to this resource
