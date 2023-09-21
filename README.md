# wgm

## Build UI

```shell
# first
cd ui/frontend
pnpm install
pnpm build
cd ../../

# normal
wails build --tags ui

# go ldflags
wails build --tags ui --ldflags='-s -w'

# with upx
wails build --tags ui --ldflags='-s -w' -upx -upxflags "--best --lzma"
```

## Build cli

```shell
go build --tags cli
```
