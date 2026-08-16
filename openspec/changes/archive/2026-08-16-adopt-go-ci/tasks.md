# Tasks — adopt-go-ci

- [x] From the meta session, run `/alemax:update-skills` so the class-M set (incl. `ci.yml`) is staged for this repo.
- [x] In this repo's session, run `/alemax:complete-update` to apply the update branch onto the working branch.
- [x] Confirm `.github/workflows/ci.yml` present and its jobs gate on `go.mod`.
- [ ] Trial push the branch; confirm `go-build`/`go-vet`/`go-test`/`golangci-lint` are green.
- [x] Confirm the rest of class-M landed: `.editorconfig`, `.gitattributes`, `.github/*`, `dependabot.yml`, `.pre-commit-config.yaml`, `bin/set-secret.sh`.
