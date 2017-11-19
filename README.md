# VcashRpcGo

Go Library for Vcash rpc commands

Need running Vcash daemon, if you don't know what's it, do not use this library

Usage example:

`response := rpc_getinfo()`

## Create standalone exe files
### i386
```
env GOOS=windows GOARCH=386 go build -o vcashRpcGo.exe vcashrpcgo.go
```

### amd64
```
env GOOS=windows GOARCH=amd64 go build -o vcashRpcGo_amd64.exe vcashrpcgo.go
```
