# Building a Talos cluster

An end-to-end walkthrough for a Talos-backed cluster on Xen Orchestra.

!!! note "Draft — not written yet"

    Planned contents:

    - The `templates/talos/base/` layout and why bases are never edited
    - Creating an overlay with real UUIDs, the VIP, and the machine config
    - The four patch files (control plane template, worker template, control
      plane, XO cluster)
    - `kubectl apply -k templates/talos/overlays/<env>/`
    - Watching the machines come up, and the maintenance-mode trap
