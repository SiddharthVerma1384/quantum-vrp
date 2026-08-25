## 📌 Description
Please include a summary of the changes and the motivation/context behind them. Mention any related issues or database migrations.

---

## 🛠️ Type of Change
- [ ] `feat`: A new feature
- [ ] `fix`: A bug fix
- [ ] `docs`: Documentation-only changes
- [ ] `refactor`: Code refactoring (no functional changes)
- [ ] `chore`: Tooling, dependencies, or maintenance
- [ ] `test`: Adding/updating tests
- [ ] `perf`: Performance optimizations

---

## 🧪 Verification & Testing
Describe how you tested your changes. Include any commands run.

- [ ] Relocated tests ran and passed (`go test -v ./...`)
- [ ] Static checks and formatters ran (`go fmt` / `go vet` / `golangci-lint`)
- [ ] Project compiles clean (`go build`)

---

## 📋 Checklist
- [ ] My commit messages conform to the **Conventional Commits** specification.
- [ ] I have updated the documentation mirror (inside `docs/`) for the modified packages.
- [ ] No raw SQL queries have been hardcoded (all queries compiled via SQLC).
- [ ] database mutations (if any) are wrapped in atomic database transactions (`pgx.Tx`).
