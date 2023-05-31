package loader

import (
	"auto_config/reader"
	"auto_config/source"
	"errors"
	"fmt"
	"time"
)

type memory struct {
	sets     []*source.ChangeSet
	sources  []source.Source
	options  Options
	SnapShot SnapShot
}

// 文件监视器，文件更新时回写内存
func (m *memory) watch(idx int, s source.Source) {
	watcher := func(idx int, s source.Watcher) error {
		for {
			cs, err := s.Next()
			if err != nil {
				return err
			}
			m.sets[idx] = cs
			fmt.Println(string(cs.Data))
			//m.SnapShot = SnapShot{Version: time.Now(), LastChange: cs}
			err = m.reload()
			if err != nil {
				return err
			}

		}

	}
	sourceWatcher, err := s.Watcher()
	if err != nil {
		panic("Create file watcher failed")
	}
	if err := watcher(idx, sourceWatcher); err != nil {
		panic("Watch file panic")
	}
}

func (m *memory) reload() error {

	merge, err := m.options.Reader.Merge(m.sets...)
	if err != nil {
		return err
	}
	m.SnapShot = SnapShot{
		Version:    time.Now(),
		LastChange: merge,
	}
	return nil
}

func (m *memory) Load(sources ...source.Source) error {
	var failedSource []interface{}
	for _, source := range sources {
		set, err := source.Read()
		if err != nil {
			failedSource = append(failedSource, sources)
			continue
		}
		m.sets = append(m.sets, set)
		m.sources = append(m.sources, source)
		idx := len(m.sets) - 1
		go m.watch(idx, source)
	}
	if len(failedSource) != 0 {
		return errors.New("ReadFile error")
	}
	// 首次转码
	if err := m.reload(); err != nil {
		return err
	}
	return nil
}

func (m *memory) Snapshot() SnapShot {
	return m.SnapShot
}

func NewLoader(opts ...Option) Loader {
	options := Options{
		Reader: reader.NewReader(),
	}
	for _, o := range opts {
		o(&options)
	}
	m := &memory{
		options: options,
	}
	return m
}
