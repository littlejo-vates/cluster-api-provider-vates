# Cluster templates

The repository ships ready-made cluster templates under `templates/`, split by
bootstrap flow:

```
templates/
├── kubeadm/
│   ├── base/          # ClusterClass + machine templates (placeholders)
│   ├── overlays/      # per-environment values (create your own)
│   ├── clusterctl/    # template for `clusterctl generate cluster`
│   └── packer/        # AlmaLinux cloud image builder
├── talos/
│   ├── base/          # Talos templates (placeholders)
│   └── overlays/      # per-environment values (create your own)
└── README.md
```

!!! note "Draft — not written yet"

    Planned contents:

    - The `base/` + `overlays/` convention, and why `base/` is never edited
    - The `clusterctl/` template and how `clusterctl generate cluster` renders it
    - The default `cluster-template.yaml` shipped in `dist/`
    - A worked overlay example for each flow

In the meantime, the raw material lives in `templates/README.md` and
`templates/kubeadm/README.md` / `templates/talos/README.md`.
