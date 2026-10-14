# Releasing

How to publish a release that users can install with `clusterctl`, without
cloning the repository.

!!! note "Draft — to be rewritten from `RELEASING.md`"

    This page will absorb the existing `RELEASING.md`, rewritten in the same
    didactic tone as the rest of the documentation. Planned contents:

    - What a `clusterctl` provider release must contain
      (`infrastructure-components.yaml`, `metadata.yaml`, `cluster-template.yaml`)
    - Building and pushing the controller image
    - Generating the release artifacts
    - Creating the GitHub release, and how the changelog is generated from
      conventional-commit labels
    - The contract version and `metadata.yaml`
