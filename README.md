# wgm

Native WireGuard manager built with Go and MyGo's native UI.

## Commands

```shell
# Develop with live rebuilds
go run github.com/egoist/mygo/cmd/mygo@v0.7.3 dev

# Package for the current platform
go run github.com/egoist/mygo/cmd/mygo@v0.7.3 build

# Test and build
go test ./...
go build ./cmd/wgm
```

MyGo reads the app, platform, and icon settings from [mygo.json](./mygo.json).
