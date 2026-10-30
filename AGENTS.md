# Development guidelines

## Architecture

- Keep the MyGo startup in `cmd/wgm`; put presentation and interaction logic in `internal/ui`.
- Keep persistence adapters in `backend` and domain/database behavior in `common`. UI code should call the adapter, not build SQL or reach into database globals.
- Preserve user workflows and existing service semantics when changing screens or validation.

## Implementation

- Prefer small, direct Go code. Add a helper only when it removes repeated logic or gives a clear boundary; avoid generic wrappers and nil/fallback handling for states the caller already guarantees.
- Keep UI state with the view that owns it. Use MyGo widgets and `Context` APIs rather than retaining per-frame `Element` values.
- Use the WireGuard red accent for primary actions. Keep destructive/error emphasis amber so it remains distinct from the brand color, and preserve system light/dark adaptation.
- Treat validation and user-facing messages as part of the workflow; don't silently coerce invalid form values to zero.
- Keep platform metadata in `mygo.json`. The top-level `icon` is a PNG source; use `macos.icon` for a custom ICNS. Do not add unsupported MyGo configuration keys.

## Tests and checks

- Add short tests for pure validation and configuration-building logic. Use MyGo's headless `ui.NewTester` only for important interactions that benefit from input-level verification.
- Keep tests independent of a real database, window server, or network where practical.
- Before finishing, run `gofmt`, `go test ./...`, `go vet ./...`, and `go build ./cmd/wgm` for Go changes.
