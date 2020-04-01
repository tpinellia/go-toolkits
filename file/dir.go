package file

import "os"

//CreateDirIfNotExist 创建目录
func CreateDirIfNotExist(fp string) error {
	return os.MkdirAll(fp, os.ModePerm)
}
