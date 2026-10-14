# Talos prerequisites

Talos needs a VM template prepared very specifically. Getting this wrong is the
usual reason a Talos machine silently stays in maintenance mode.

!!! note "Draft — not written yet"

    Planned contents:

    - The `nocloud` platform build (Talos image factory)
    - The `siderolabs/xen-guest-agent` extension, and why the guest agent matters
    - The template must be **never booted** and set to `viridian: false` in XO
    - Pinning the Talos version in
      `TalosControlPlane.spec.controlPlaneConfig.controlplane.talosVersion`
    - Why the version must match the image, and what happens if it does not
