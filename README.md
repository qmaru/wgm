# wgm

Build Binary

```shell
# normal
wails build --tags ui

# go ldflags
wails build --tags ui --ldflags='-s -w'

# with upx
wails build --tags ui --ldflags='-s -w' -upx -upxflags "--best --lzma"
```
