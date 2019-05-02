package classpath

import (
	"os"
	"path/filepath"
)

type Classpath struct {
	bootClasspath Entry
	extClasspath  Entry
	userClasspath Entry
}

func Parse(jreOption string, cpOption string) *Classpath {
	classpath := &Classpath{}
	classpath.parseBootAndExtClassPath(jreOption)
	classpath.parseUserClasspath(cpOption)

	return classpath
}

func (classpath *Classpath) parseUserClasspath (cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	classpath.userClasspath = newEntry(cpOption)
}

func (classpath *Classpath) parseBootAndExtClassPath(jreOption string) {
	jreDir := classpath.getJreDir(jreOption)

	// /jre/lib
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	classpath.bootClasspath = newWildcardEntry(jreLibPath)

	// /jre/ext/*
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	classpath.extClasspath = newWildcardEntry(jreExtPath)
}

func (classpath *Classpath) getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}
	if exists("./jre") {
		return "./jre"
	}
	if javaHome := os.Getenv("JAVA_HOME"); javaHome != "" {
		return filepath.Join(javaHome, "jre")
	}
	panic("Can't find jre folder.")
}


func (classpath *Classpath) String() string {
	return classpath.userClasspath.String()
}

func (classpath *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className += ".class"
	if data, entry, err := classpath.bootClasspath.ReadClass(className); err == nil {
		return data, entry, err
	}

	if data, entry, err := classpath.extClasspath.ReadClass(className); err == nil {
		return data, entry, err
	}

	return classpath.userClasspath.ReadClass(className)
}

func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
