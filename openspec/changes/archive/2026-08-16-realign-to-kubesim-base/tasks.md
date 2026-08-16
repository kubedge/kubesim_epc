# Tasks — realign-to-kubesim-base

- [x] Wait for `kubesim_base` to publish the new tag (its multimodule-tag-realign change).
- [x] `go get` each kubesim_base sub-module this repo requires (`config` and/or `connected`/`grpc/go`) at the new `<tag>`.
- [x] `go mod tidy`; confirm every kubesim_base require moved to `<tag>`.
- [x] `go build ./... && go vet ./... && go test ./... -race` green.
