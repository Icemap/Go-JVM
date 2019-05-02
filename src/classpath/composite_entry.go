package classpath

import (
	"errors"
	"strings"
)

type CompositeEntry []Entry

func newCompositeEntry(pathList string) CompositeEntry {
	var compositeEntry []Entry
	for _, path := range strings.Split(pathList, pathListSep) {
		entry := newEntry(path)
		compositeEntry = append(compositeEntry, entry)
	}
	return compositeEntry
}

func (compositeEntry CompositeEntry) ReadClass(className string) ([]byte, Entry, error) {
	for _, entry := range compositeEntry {
		data, entryItem, err := entry.ReadClass(className)
		if err == nil {
			return data, entryItem, nil
		}
	}
	return nil, nil, errors.New("Class not found: " + className)
}

func (compositeEntry CompositeEntry) String() string {
	result := ""
	for i, entry := range compositeEntry {
		if i != 0 {
			result += pathListSep
		}
		result += entry.String()
	}
	return result
}

