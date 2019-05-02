package classpath

import (
	"os"
	"strings"
)

const pathListSep = string(os.PathListSeparator)

type Entry interface {
	readClass(className string) ([]byte, Entry, error)
	String() string
}

/**
 *	新建一个类路径，用以加载Class文件
 *	类路径有四种方式传入：文件夹路径、直接的Jar包或Zip包路径、联合路径与通配符路径
 *	分别对应四种Entry接口的实现
 */
func newEntry(path string) Entry {
	if strings.Contains(path, pathListSep) {
		// 存在分隔符，联合路径
		return newCompositeEntry(path)
	} else if strings.HasSuffix(path, "*") {
		// 通配符在结尾，通配符路径，其实也是联合路径的一种
		return newWildcardEntry(path)
	} else if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".zip") ||
		strings.HasSuffix(path, ".JAR") || strings.HasSuffix(path, ".ZIP") {
		// zip文件或者jar文件在结尾，是直接的包路径
		return newZipEntry(path)
	} else {
		// 什么都没有的，是文件夹路径
		return newDirEntry(path)
	}
}