package loader

import (
	"auto_config/source/file"
	"fmt"
	"testing"
)

func TestMemory_Load(t *testing.T) {
	source := file.NewSource(file.WithPath("./test/test.yaml"))
	loader := NewLoader()
	if err := loader.Load(source); err != nil {
		t.Error(err)
	}
	w := loader.Watcher()
	for {
		sp, err := w.Next()
		fmt.Println(sp, err)
		if err != nil {
			return
		}

		w.Stop()
	}
	//select {}

}
