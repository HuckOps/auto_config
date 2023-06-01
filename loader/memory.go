package loader

import (
	"auto_config/reader"
	"auto_config/source"
	"errors"
	"sync"
	"time"
)

type memory struct {
	sets     []*source.ChangeSet
	sources  []source.Source
	options  Options
	SnapShot SnapShot
	watcher  *watcher
	sync.RWMutex
}
type watcher struct {
	updates chan SnapShot
}

// 文件监视器，文件更新时回写内存
func (m *memory) watch(idx int, s source.Source) {
	watcher := func(idx int, s source.Watcher) error {
		for {
			cs, err := s.Next()
			m.Lock()
			if err != nil {
				m.Unlock()
				return err
			}
			m.sets[idx] = cs
			//m.SnapShot = SnapShot{Version: time.Now(), LastChange: cs}
			err = m.reload()
			//m.watcher.updates <- m.SnapShot
			m.Unlock()
			m.update()
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
		panic(err)
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
		//写锁，防止其他线程串入数据
		m.Lock()
		m.sets = append(m.sets, set)
		m.sources = append(m.sources, source)
		idx := len(m.sets) - 1
		m.Unlock()
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

func (m *memory) update() {
	m.watcher.updates <- m.SnapShot
	//w.update <- snapshot
	//m.watcher m.SnapShot
}

func (m *memory) Watcher() Watcher {
	w := &watcher{
		updates: make(chan SnapShot, 1),
	}
	m.watcher = w
	return w
}

func (w *watcher) Next() (*SnapShot, error) {
	for {
		select {
		case u := <-w.updates:
			return &u, nil
		}
	}
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
