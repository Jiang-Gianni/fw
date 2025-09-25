# fw
File Watcher

This tool generates a go template file like [this](./cmd/fw/fw.go) which can be run as Go program to listen to file updates and execute operations.

## Why?

There are already many tools that would do the same job (examples [**air**](https://github.com/air-verse/air), [**task**](https://github.com/go-task/task)) but they require creating some configuration files.

I really didn't want to read all their documentation, their settings etc... so I decided to ~~invest 1000X the time it would take me to use one of the tools above and~~ build my own tool with a template Go program that listen to file change events and run some operations.

The other tools I mentioned above are also written in Go and use [fsnotify](https://github.com/fsnotify/fsnotify), which is just what is really needed for my needs.

## Installation

```bash
go install github.com/Jiang-Gianni/fw@latest
```

and then run the following inside the desired directory (requires a `go.mod` file -> a Go project):

```bash
fw
```

which will create (if not already present) the directory `watch/fw` and the file `fw.go` in that directory.

To specify the destination directory and filename:

```bash
# This creates ./myDir/mySubDir/myWatcher.go
fw -d myDir/mySubDir -f myWatcher.go
```

Update the generated file `watch/fw/fw.go` with what you need, resolve the dependencies with `go mod tidy` and just run the program.

The generated Go template has some default example functions to live reloading/generating:
- Go application
- [sqlc](https://github.com/sqlc-dev/sqlc) (**sqlc** CLI tool needed)
- [templ](https://github.com/a-h/templ) (**templ** CLI tool needed)
- [esbuild](https://github.com/evanw/esbuild)
- [brotli](https://github.com/andybalholm/brotli)
- [buf](https://github.com/bufbuild/buf)
- [markdown](https://github.com/yuin/goldmark)
- [d2](https://github.com/terrastruct/d2) (**d2** CLI tool needed)
- [webp](https://github.com/chai2010/webp)