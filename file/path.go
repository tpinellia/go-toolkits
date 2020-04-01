package file

import (
	"os"
	"path"
	"path/filepath"
)

//GetSelfAbsolutePath 获取当前可执行文件的绝对路径
func GetSelfAbsolutePath() (string, error) {
	path, err := filepath.Abs(os.Args[0])
	return path, err
}

//GetSelfAbsoluteDir 获取当前可执行文件所在目录的绝对路径
func GetSelfAbsoluteDir() (string, error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	return dir, err
}

//GetFilePath 获取fp的绝对路径，fp为文件或路径
func GetFilePath(fp string) (string, error) {
	if path.IsAbs(fp) {
		return fp, nil
	}
	wd, err := os.Getwd()
	return path.Join(wd, fp), err
}

//GetFileDir 获取fp的上级目录，fp为文件或路径.该函数只是简单的进行路径处理，例如:fp为"/tmp",返回"/";fp为"../logger/logger.go",返回"../logger";fp为"logger.go"，返回".";
func GetFileDir(fp string) string {
	return path.Dir(fp)
}

//GetCurrentDir 获取当前目录，返回绝对路径
func GetCurrentDir() (string, error) {
	return os.Getwd()
}
