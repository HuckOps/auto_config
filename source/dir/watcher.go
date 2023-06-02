package dir

import (
	"errors"
	"github.com/fsnotify/fsnotify"
)

type dir struct {
	path string
	fw   *fsnotify.Watcher
}

type Watcher interface {
	Next() (string, error)
}

func (w *dir) Next() (string, error) {
	for {
		select {
		case event, _ := <-w.fw.Events:
			//if event.Op == fsnotify.Create {
			return event.Name, nil
			//}
		case <-w.fw.Errors:
			return "", errors.New("Watch dir field")

		}
	}

}

func NewWatcher(path string) (Watcher, error) {
	fswatcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}
	fswatcher.Add(path)
	return &dir{
		path: path,
		fw:   fswatcher,
	}, nil
}
