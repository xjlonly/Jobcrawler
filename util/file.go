package util

import "os"

//检查指定的文件或目录是否存在 存在则返回true 不存在则返回flase
func IsExist(filename string) bool {
	_,err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}
