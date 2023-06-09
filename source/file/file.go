package file

import (
	"github.com/huckops/auto_config/source"
	"io/ioutil"
	"os"
	"strings"
)

type file struct {
	path string
	opts source.Options
}

func FileFormat(path string) string {
	filename := strings.Split(path, "/")
	suffix := filename[len(filename)-1]
	parts := strings.Split(suffix, ".")
	return parts[len(parts)-1]
}

func (f *file) Read() (*source.ChangeSet, error) {
	file, err := os.Open(f.path)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	info, err := file.Stat()
	if err != nil {
		return nil, err
	}
	changeSet := &source.ChangeSet{
		Data:      bytes,
		Format:    FileFormat(f.path),
		Timestamp: info.ModTime(),
	}
	changeSet.CheckSum = changeSet.Sum()
	return changeSet, nil
}

func (f *file) Write(*source.ChangeSet) (*source.ChangeSet, error) {
	return &source.ChangeSet{}, nil
}

func (f *file) Watcher() (source.Watcher, error) {
	return NewWatcher(f)
}

func (f *file) Path() string {
	return f.path
}

func NewSource(opts ...source.Option) source.Source {
	options := source.NewOptions(opts...)
	f, _ := options.Context.Value(fileKeyPath{}).(string)

	return &file{path: f, opts: options}
}
