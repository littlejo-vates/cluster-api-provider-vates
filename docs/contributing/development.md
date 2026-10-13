# Development

Building the provider from source and running it against your own management
cluster.

!!! note "Draft — to be rewritten from `DEVELOPMENT.md`"

    This page will absorb the existing `DEVELOPMENT.md`, rewritten in the same
    didactic tone as the rest of the documentation. Planned contents:

    - Prerequisites and the DevContainer
    - One-time management cluster setup
    - The `Makefile.dev` build/push/deploy loop (kind and k3s)
    - Fast iteration with `make run`
    - `make manifests` and `make generate` after type changes
    - The Helm chart workflow (`make helm-update-chart`, the `--force` trap)
    - Debugging and undeploying
