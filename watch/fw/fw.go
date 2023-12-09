// Code generated by fw from github.com/Jiang-Gianni/fw
// Any update will be overridden after regenerating the file.

package main

import (
	"log"
	"os"
	"os/exec"
	"regexp"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
)

// File Watcher struct.
type FW struct {
	// Directories to watch. Subdirectories are automatically included.
	directories []string

	// Files to watch if not included in the above directories.
	files []string

	// Regex to be matched with the updated file name.
	regexMatch string

	// Some code editor triggers a file update twice per save.
	// Set this ticker to avoid multiple execution at the same time.
	// Consider it as a rate limiter.
	ticker *time.Ticker

	// Function to run when a file update is triggered and the filename matches regexMatch.
	run WatchFunc

	// Other fields that can be used in the run function.
	signal chan (struct{})
}

type WatchFunc func(filename string, fw *FW)

var fws = []FW{
	{
		directories: []string{"./"},
		files:       []string{"main.go"},
		regexMatch:  ".go$",
		ticker:      time.NewTicker(time.Second),
		signal:      make(chan struct{}),
		run:         GoWatch(),
	},
}

func main() {
	for _, fw := range fws {
		go Watch(fw)
	}

	var block chan struct{}

	<-block
}

// Function to run on go files updates.
func GoWatch() WatchFunc {
	var init = sync.Once{}
	return func(filename string, fw *FW) {
		// This goroutine blocks on fw.signal read
		// If main.go starts a server, and a file is updated (<-fw.signal)
		// kill the process using the package github.com/pbar1/pkill-go
		go init.Do(
			func() {
				for {
					<-fw.signal

					// _, err := pkill.Pkill("main", os.Kill)
					// if err != nil {
					// 	log.Println(err)
					// }

					cmd := exec.Command("go", "run", "main.go")
					go RunCmd(cmd, true)
				}
			},
		)

		// Run some other stuff (example: code formatters)
		// before sending a signal to restart
		log.Println("File updated: ", filename)
		fw.signal <- struct{}{}
	}
}

func Watch(fw FW) {
	rm, err := regexp.Compile(fw.regexMatch)
	if err != nil {
		log.Fatal(err)
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	defer fw.ticker.Stop()

	for _, file := range fw.files {
		err = watcher.Add(file)
		if err != nil {
			log.Fatal(err)
		}
	}

	for _, dir := range fw.directories {
		err = AddDir(dir, watcher)
		if err != nil {
			log.Fatal(err)
		}
	}

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}

			if rm.MatchString(event.Name) {
				select {
				case <-fw.ticker.C:
					fw.run(event.Name, &fw)
				default:
				}
			}

		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}

			log.Println("error:", err)
		}
	}
}

func RunCmd(cmd *exec.Cmd, sameStdout bool) {
	if sameStdout {
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
	}

	if err := cmd.Run(); err != nil {
		log.Println(err)
	}
}

func AddDir(dir string, watcher *fsnotify.Watcher) error {
	err := watcher.Add(dir)
	log.Println(dir)
	if err != nil {
		log.Fatal(err)
	}
	entries, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range entries {
		if entry.IsDir() && entry.Name() != ".git" {
			var entryName string
			if dir == "./" {
				entryName = "./" + entry.Name()
			} else {
				entryName = dir + "/" + entry.Name()
			}

			AddDir(entryName, watcher)
		}
	}
	return nil
}
