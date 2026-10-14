# kube-vip

kube-vip provides the floating control plane endpoint (`CP_HOST`) that all
control plane nodes advertise, so the API server survives losing a node.

With vates-capi, kube-vip is *injected* into the control plane machines'
cloud-init when the cluster asks for it — the provider does not run kube-vip
itself.

!!! note "Draft — not written yet"

    Planned contents:

    - What a floating control plane endpoint is, and why you need one
    - Enabling it: `XOCluster.spec.controlPlaneLB: kube-vip`
    - The matching declaration on the machines: `XOMachine.spec.bootstrap.format: cloudinit`
    - How the injection works (a generated static pod manifest)
    - Verifying the VIP has been claimed
    - Troubleshooting: the VIP never coming up
