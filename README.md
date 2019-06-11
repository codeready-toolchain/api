# api
ToolChain API

# Building
CodeReady ToolChain API is built using [Go modules][go_modules].  Set the following environment variable in your CLI before building:

```sh
$ export GO111MODULE=on
```

Execute the go build command to build the API:

```sh
$ go build ./...
```

[go_modules]: https://github.com/golang/go/wiki/Modules