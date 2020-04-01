package file

import (
	"os"
	"path"
)

//WriteBytesToFile 将[]byte写入文件
func WriteBytesToFile(filePath string, b []byte) (int, error) {
	os.MkdirAll(path.Dir(filePath), os.ModePerm)
	fw, err := os.Create(filePath)
	if err != nil {
		return 0, err
	}
	defer fw.Close()
	return fw.Write(b)
}

//WriteStringToFile 将字符串写入文件
func WriteStringToFile(filePath string, s string) (int, error) {
	return WriteBytesToFile(filePath, []byte(s))
}
