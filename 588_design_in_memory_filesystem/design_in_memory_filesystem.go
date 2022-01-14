package design_in_memory_filesystem

import (
	"sort"
	"strings"
)

type Entry struct {
	file       bool
	name       string
	content    string
	subEntries map[string]*Entry
}

type FileSystem struct {
	root *Entry
}

func Constructor() FileSystem {
	return FileSystem{
		root: &Entry{
			name:       "root",
			subEntries: make(map[string]*Entry, 0),
		},
	}
}

func (this *FileSystem) get(path string) *Entry {
	entry := this.root

	for _, name := range strings.Split(path, "/")[1:] {
		if name == "" {
			return entry
		}

		if v, ok := entry.subEntries[name]; ok {
			entry = v
		} else {
			return nil
		}
	}

	return entry
}

func (this *FileSystem) Ls(path string) []string {
	e := this.get(path)

	if e.file {
		return []string{e.name}
	}

	names := make([]string, 0, len(e.subEntries))
	for _, entry := range e.subEntries {
		names = append(names, entry.name)
	}

	sort.Strings(names)

	return names
}

func (this *FileSystem) Mkdir(path string) {
	entry := this.root

	for _, name := range strings.Split(path, "/")[1:] {
		if name == "" {
			break
		}

		if v, ok := entry.subEntries[name]; ok {
			entry = v
		} else {
			newEntry := &Entry{
				name:       name,
				subEntries: make(map[string]*Entry),
			}

			entry.subEntries[name] = newEntry
			entry = newEntry
		}
	}
}

func (this *FileSystem) AddContentToFile(filePath string, content string) {
	entry := this.get(filePath)

	if entry != nil {
		entry.content += content
		return
	}

	li := strings.LastIndex(filePath, "/")
	dir, file := filePath[:li], filePath[li+1:]

	this.Mkdir(dir)
	folder := this.get(dir)

	folder.subEntries[file] = &Entry{
		file:    true,
		name:    file,
		content: content,
	}
}

func (this *FileSystem) ReadContentFromFile(filePath string) string {
	entry := this.get(filePath)

	if entry != nil && entry.file {
		return entry.content
	} else {
		return ""
	}
}
