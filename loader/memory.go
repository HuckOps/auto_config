package loader

import (
	"auto_config/common"
	"auto_config/reader"
	"auto_config/source"
	"auto_config/source/dir"
	"errors"
	"fmt"
	"strings"
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
	exit chan bool
}
type watcher struct {
	updates chan SnapShot
	exit    chan bool
}

func (m *memory) dirWatcher(path string) {
	w, _ := dir.NewWatcher(path)
	for {
		fmt.Println("test")
		file, _ := w.Next()
		fmt.Println(file)
	}
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
	//loaderWatcher := m.Watcher()
	sourceWatcher, err := s.Watcher()
	if err != nil {
		panic("Create file watcher failed")
	}
	lwBreak := make(chan bool)
	go func() {
		select {
		case <-lwBreak:
		}
		sourceWatcher.Stop()
	}()
	// 监听源监听器
	if err := watcher(idx, sourceWatcher); err != nil {
		fmt.Println(err)
		time.Sleep(time.Second)
	}
	// 监听异常时关闭通道
	close(lwBreak)
	//// 关闭装载器
	//select {
	//case <-m.exit:
	//	return
	//}
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

func GetDir(path string) string {
	parts := strings.Split(path, "/")
	dirnamePart := parts[0 : len(parts)-2]
	return strings.Join(dirnamePart, "/")
}

func (m *memory) Load(sources ...source.Source) error {
	var failedSource []interface{}
	var confDirs []string
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
		confDirs = append(confDirs, GetDir(source.Path()))
		m.Unlock()
		go m.watch(idx, source)

	}
	confDirs = common.RemoveRepeatedElementAndEmpty(confDirs)
	for _, dir := range confDirs {
		go m.dirWatcher(dir)
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
	select {
	case m.watcher.updates <- m.SnapShot:

	}

	//w.update <- snapshot
	//m.watcher m.SnapShot
}

func (m *memory) Watcher() Watcher {
	w := &watcher{
		updates: make(chan SnapShot, 1),
		exit:    make(chan bool),
	}
	m.watcher = w
	go func() {

	}()
	return w
}

func (w *watcher) Next() (*SnapShot, error) {
	for {
		select {
		case e := <-w.exit:
			fmt.Println(e)
			return nil, errors.New("watcher stop")
		case u := <-w.updates:
			return &u, nil
		}
	}
}

func (w *watcher) Stop() error {
	select {
	//case <-w.exit:
	default:
		close(w.exit)
		close(w.updates)
	}
	return nil
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
