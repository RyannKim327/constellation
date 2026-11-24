package main

import (
	"io/fs"
	"path/filepath"
	"strings"
	"time"
)

const (
	pollInterval = 400 * time.Millisecond
	settleDelay  = 250 * time.Millisecond
)

type fileState struct {
	modTime time.Time
	size    int64
}

type watcher struct {
	root string
	last map[string]fileState
}

func newWatcher(root string) (*watcher, error) {
	state, err := scan(root)
	if err != nil {
		return nil, err
	}
	return &watcher{root: root, last: state}, nil
}

func (w *watcher) changed() bool {
	time.Sleep(pollInterval)

	current, err := scan(w.root)
	if err != nil || !different(w.last, current) {
		return false
	}

	time.Sleep(settleDelay)
	current, err = scan(w.root)
	if err != nil {
		return false
	}

	changed := different(w.last, current)
	w.last = current
	return changed
}

func scan(root string) (map[string]fileState, error) {
	state := make(map[string]fileState)

	err := filepath.WalkDir(root, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}

		name := entry.Name()
		if entry.IsDir() {
			if strings.HasPrefix(name, ".") || name == "vendor" || name == "node_modules" || name == "tmp" {
				return filepath.SkipDir
			}
			return nil
		}

		if strings.HasPrefix(name, ".") || strings.HasSuffix(name, "~") || strings.HasSuffix(name, ".swp") || strings.HasSuffix(name, ".tmp") {
			return nil
		}

		ext := strings.ToLower(filepath.Ext(name))
		if ext != ".go" && ext != ".css" && ext != ".md" {
			return nil
		}

		info, err := entry.Info()
		if err != nil {
			return nil
		}
		state[path] = fileState{modTime: info.ModTime(), size: info.Size()}
		return nil
	})

	return state, err
}

func different(previous, current map[string]fileState) bool {
	if len(previous) != len(current) {
		return true
	}
	for path, state := range current {
		old, ok := previous[path]
		if !ok || !state.modTime.Equal(old.modTime) || state.size != old.size {
			return true
		}
	}
	return false
}
