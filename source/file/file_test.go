package file

import (
	"fmt"
	"testing"
)

func TestFile(t *testing.T) {
	source := NewSource(WithPath("../test/test.yaml"))
	changeSet, err := source.Read()
	if err != nil {
		t.Error(changeSet)
	}
	fmt.Println(changeSet)
}

func TestWatcher(t *testing.T) {
	source := NewSource(WithPath("../test/test.yaml"))
	watcher, _ := source.Watcher()
	go func() {
		for {
			set, err := watcher.Next()
			if err != nil {
				t.Error(err)
			}
			fmt.Println(set)
		}
	}()
	select {}
}
