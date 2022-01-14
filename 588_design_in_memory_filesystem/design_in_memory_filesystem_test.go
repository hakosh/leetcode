package design_in_memory_filesystem

import (
	"reflect"
	"testing"
)

func TestFilesystem(t *testing.T) {
	fs := Constructor()

	if res := fs.Ls("/"); !reflect.DeepEqual(res, []string{}) {
		t.Errorf("For ls(/) expected %v got %v", []string{}, res)
	}

	fs.Mkdir("/a/b/c")
	fs.Mkdir("/a/b/c")

	fs.AddContentToFile("/a/b/c/d", "hello")
	fs.AddContentToFile("/a/b/e", "eeeee")

	if res := fs.Ls("/"); !reflect.DeepEqual(res, []string{"a"}) {
		t.Errorf("For ls(/) expected %v, got %v", []string{"a"}, res)
	}

	if res := fs.Ls("/a/b"); !reflect.DeepEqual(res, []string{"c", "e"}) {
		t.Errorf("For ls(/) expected %v, got %v", []string{"c", "e"}, res)
	}

	if res := fs.ReadContentFromFile("/a/b/c/d"); res != "hello" {
		t.Errorf("For readContentFromFile(/a/b/c/d) expected %v, got %v", "hello", res)
	}
}
