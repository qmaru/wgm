# wgm

Build Binary

```node
wails build --tags ui --ldflags='-s -w' -upx -upxflags "--best --lzma"

# go ldflags
wails build --tags ui --ldflags='-s -w'

# with upx
wails build --tags ui --ldflags='-s -w' -upx -upxflags "--best --lzma"
```
