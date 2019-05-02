package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

type ZipEntry struct {
	absPath string
}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}

	return &ZipEntry{absPath}
}

func (zipEntry *ZipEntry) String() string{
	return zipEntry.absPath
}

func (zipEntry *ZipEntry) ReadClass(className string) ([]byte, Entry, error) {
	// 遍历Zip或Jar内部，寻找class文件
	reader, err := zip.OpenReader(zipEntry.absPath)
	if err != nil {
		panic(err)
	}

	defer reader.Close()

	for _, innerFile := range reader.File {
		if innerFile.Name == className {
			return readZipInnerFile(innerFile, zipEntry)
		}
	}

	return nil, nil, errors.New("Class not found: " + className)
}

// 不在循环内部调用defer，规避编译器风险
func readZipInnerFile (innerFile *zip.File, zipEntry *ZipEntry) ([]byte, Entry, error) {
	innerReader, err := innerFile.Open()
	if err != nil {
		return nil, nil, err
	}
	defer innerReader.Close()

	data, err := ioutil.ReadAll(innerReader)
	if err != nil {
		return nil, zipEntry, err
	}

	return data, zipEntry, nil
}