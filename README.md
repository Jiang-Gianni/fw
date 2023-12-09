# fw
File Watcher

This tool generates a go template file like [this](./watch/fw/fw.go) which can be run to watch and execute a command after any file update.


## Why?

There are many tools that would do the same job but they require learning and/or setting some configuration files or similar, so I decided to reinvent the wheel and to build one of my own.

With the generated file it is possible to define the directories/files to watch and the commands to run on any file update.

It uses [fsnotify](https://github.com/fsnotify/fsnotify).



## Installation

```bash
go install github.com/Jiang-Gianni/fw@latest
```

and then run the following inside the desired directory:

```bash
fw
```

which will create (if not already present) the directory `watch/fw` and the file `fw.go` in that directory.

To specify the destination directory and filename:

```bash
# This creates ./myDir/mySubDir/myWatcher.go
fw -d myDir/mySubDir -f myWatcher.go
```

After resolving the [fsnotify](https://github.com/fsnotify/fsnotify) dependency (`go mod tidy`), update `watch/fw/fw.go` and run:

```bash
go run watch/fw/fw.go
```



## How to setup

### Example

Inside the generated file, update the variable `fws` by adding a value for any filetype / directory / file to watch.

In the generated template:

```go
	{
		directories: []string{"./"},
		files:       []string{"main.go"},
		regexMatch:  ".go$",
		ticker:      time.NewTicker(time.Second),
		signal:      make(chan struct{}),
		run:         GoWatch(),
	},
```

will watch all the subdirectories and the `main.go` files. The function definede in the `run` field is executed when a `.go` file is updated (`regexMatch`) at a rate of maximum one per second (`ticker` field, the reason for this is that some code editors trigger multiple file updates at once).


### Another example

Another example: if you have tools like [sqlc](https://github.com/sqlc-dev/sqlc) installed, this file watcher tool can execute `sqlc generate` by adding the following to the `fws` list (adjust directories and files watch based on your repository):


```go
	{
		directories: []string{"./sql"},
		files:       []string{},
		regexMatch:  ".sql$",
		ticker:      time.NewTicker(time.Second),
		signal:      make(chan struct{}),
		run: func(filename string, fw *FW) {
			log.Println("SQL file update: ", filename)
			cmd := exec.Command("sqlc", "generate")
			RunCmd(cmd, true)
		},
	},
```

This will generate/update the go files which in turn will trigger the go file watcher (if present).
