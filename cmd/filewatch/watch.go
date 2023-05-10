package main

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/reagancn/filewatch/pkg/utils"
)

func Watch(paths []string) {

	if len(paths) < 1 {
		utils.Exit("must specify at least one path to watch", usage)
	}

	// Create a new watcher.
	w, err := fsnotify.NewWatcher()
	if err != nil {
		utils.Exit("creating a new watcher: %s", usage, err)
	}
	defer w.Close()

	// Start listening for events.
	go watchLoop(w)

	// Add all paths from the commandline.
	for _, p := range paths {
		err = w.Add(p)
		if err != nil {
			utils.Exit("%q: %s", usage, p, err)
		}
	}

	log.Printf("ready; press ^C to exit")
	<-make(chan struct{}) // Block forever
}

func watchLoop(w *fsnotify.Watcher) {
	i := 0
	for {
		select {
		// Read from Errors.
		case err, ok := <-w.Errors:
			if !ok { // Channel was closed (i.e. Watcher.Close() was called).
				return
			}
			log.Printf("ERROR: %s", err)
		// Read from Events.
		case e, ok := <-w.Events:
			if !ok { // Channel was closed (i.e. Watcher.Close() was called).
				return
			}

			// Just print the event nicely aligned, and keep track how many
			// events we've seen.
			i++
			log.Printf("%3d %s", i, e)
		}
	}
}
