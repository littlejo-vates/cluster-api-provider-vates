# SSH keys

Injecting the Xen Orchestra user's SSH public keys into every VM's
`ssh_authorized_keys`, so you can log in to nodes for debugging.

!!! note "Draft — not written yet"

    Planned contents:

    - What `XOCluster.spec.injectSSHKeys` does (and that it is off by default)
    - Why it is opt-in: it modifies user-provided bootstrap data
    - The pairing with `XOMachine.spec.bootstrap.format: cloudinit`
    - Where the keys come from (the authenticated XO user profile)
    - Verifying the keys landed in the VM
