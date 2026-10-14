# kubeadm

The kubeadm flow is the default way to build clusters with vates-capi. The VMs
boot a cloud-init capable OS, and the kubeadm bootstrap provider hands them the
configuration that turns them into control plane or worker nodes.

!!! note "Draft — not written yet"

    This page will be written from `README.md` and
    `templates/kubeadm/README.md`. Planned contents:

    - The cloud-init based lifecycle, step by step
    - Choosing a machine image (AlmaLinux `from-scratch` vs pre-filled)
    - Generating a cluster from the clusterctl template
    - The `base/` + `overlays/` layout and how to create an overlay
    - Version pinning: Kubernetes and kubeadm versions
    - Building the AlmaLinux image with Packer
