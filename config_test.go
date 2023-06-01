package auto_config

import (
	"auto_config/loader"
	"auto_config/source/file"
	"fmt"
	"testing"
)

func TestConfig(t *testing.T) {
	var Dest map[string]interface{}
	//sources := []Option{
	//	WithSource(file.NewSource(file.WithPath("./test/test.yaml"))),
	//}
	//dests := []interface{}{
	//	dest,
	//}
	w := func() {
		fmt.Println(Dest)
	}
	config, err := NewConfig(WithSource(file.NewSource(file.WithPath("./test/test.yaml"))), WithEntity(&Dest), WithCallback(w))
	if err != nil {
		t.Error(err)
		panic(err)
	}
	config.Watcher()
	//for {
	//	fmt.Println(Dest)
	//}
	select {}
}

func TestWatcher(t *testing.T) {
	var i interface{}
	config, _ := NewConfig(WithSource(file.NewSource(file.WithPath("./test/test.yaml"))), WithEntity(i))
	w := config.opts.Loader.Watcher()
	go func(w loader.Watcher) {
		for {
			t, err := w.Next()
			fmt.Println(t, err)
		}
	}(w)
	select {}
}

func TestT(t *testing.T) {
	t1 := make(chan string, 1)
	go func(t chan string) {
		for {
			t <- "s"
		}
	}(t1)
	for {
		select {
		case t := <-t1:
			fmt.Println(t)
		}
	}
}
