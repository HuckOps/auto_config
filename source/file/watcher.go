package file

import (
	"auto_config/source"
	"github.com/fsnotify/fsnotify"
	"os"
)

type watcher struct {
	f  *file
	fw *fsnotify.Watcher
}

func (w *watcher) Next() (*source.ChangeSet, error) {
	select {
	case event, _ := <-w.fw.Events:
		if event.Op == fsnotify.Rename {
			// check existence of file, and add watch again
			_, err := os.Stat(event.Name)
			if err == nil || os.IsExist(err) {
				err := w.fw.Add(event.Name)
				if err != nil {
					return nil, err
				}
			}
		}
		c, err := w.f.Read()
		if err != nil {
			return nil, err
		}
		return c, nil
	case err := <-w.fw.Errors:
		return nil, err
	}
}

func NewWatcher(f *file) (source.Watcher, error) {
	fswatcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}
	fswatcher.Add(f.path)
	return &watcher{
		f:  f,
		fw: fswatcher,
	}, nil
}
