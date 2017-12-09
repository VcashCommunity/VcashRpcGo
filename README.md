# VcashRpcGo

Go library for Vcash rpc commands

### Installing
To start using VcashRpcGo, install Go 1.8 or above and run `go get`:

```sh
$ go get -u github.com/devmahno/vcashrpcgo
```

This will retrieve the library and install the `VcashRpcGo`
utility into your `$GOPATH`.


### Usage example
Need running Vcash daemon, if you don't know what's it, do not use this library


`response := RpcGetInfo()`


## Create standalone exe files
### i386
```
env GOOS=windows GOARCH=386 go build -o vcashRpcGo.exe vcashrpcgo.go
```

### amd64
```
env GOOS=windows GOARCH=amd64 go build -o vcashRpcGo_amd64.exe vcashrpcgo.go
```
