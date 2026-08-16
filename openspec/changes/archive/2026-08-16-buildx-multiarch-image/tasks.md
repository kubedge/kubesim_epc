# Tasks — buildx-multiarch-image

- [x] Confirm buildx + a running builder (colima on Apple-Silicon).
- [x] For each built image, add a buildx target `docker buildx build --platform linux/arm64 -t <image>:<tag>` (add `,linux/amd64` only if still needed).
- [x] Make each Dockerfile multi-stage (compile per TARGETOS/TARGETARCH; don't copy a prebuilt amd64 binary).
- [x] Remove the `_AMD64/_ARM64V8/_ARM32V7` image-name variables.
- [x] `docker buildx imagetools inspect` each image → manifest list incl. linux/arm64.
- [x] If deployed by an operator, update that operator's example CR to the arch-independent name.
