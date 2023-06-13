package auto_config

import (
	"fmt"
	"github.com/huckops/auto_config/loader"
	"github.com/huckops/auto_config/source/file"
	"testing"
)

func TestConfig(t *testing.T) {
	var Dest map[string]interface{}
	w := func() {
		fmt.Println(Dest)
	}
	config, err := NewConfig(WithSource(file.NewSource(file.WithPath("/home/huck/桌面/auto_config/test/test.yaml"))), WithSource(file.NewSource(file.WithPath("/home/huck/桌面/auto_config/test/test2.yaml"))), WithEntity(&Dest), WithCallback(w))
	if err != nil {
		t.Error(err)
		panic(err)
	}
	config.Watcher()
	//select {}
}

func TestWatcher(t *testing.T) {
	var i interface{}
	config, _ := NewConfig(WithSource(file.NewSource(file.WithPath("/home/huck/桌面/auto_config/test/test.yaml"))), WithEntity(i))
	w := config.opts.Loader.Watcher()
	go func(w loader.Watcher) {
		for {
			t, err := w.Next()
			fmt.Println(t, err)
		}
	}(w)
	select {}
}
