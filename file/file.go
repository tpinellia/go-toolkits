package file

import "os"

//IsFile 判断路径是否为文件,返回true时是文件，false表示路径为文件夹或路径不存在
func IsFile(fp string) bool {
	f, err := os.Stat(fp)
	return err == nil && !f.IsDir()
}

//IsExist 判断文件或路径是否存在
func IsExist(fp string) bool {
	_, err := os.Stat(fp)
	return err == nil || os.IsExist(err)
}
