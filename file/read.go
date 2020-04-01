package file

import (
	"io/ioutil"
	"strconv"
	"strings"
)

//ReadToBytes 读文件返回[]byte
func ReadToBytes(filePath string) ([]byte, error) {
	return ioutil.ReadFile(filePath)
}

//ReadToString 读文件并转化为string
func ReadToString(filePath string) (string, error) {
	if contents, err := ioutil.ReadFile(filePath); err == nil {
		return string(contents), nil
	} else {
		return "", err
	}
}

//ReadToTrimString 读文件并且移除前后空格
func ReadToTrimString(filePath string) (string, error) {
	str, err := ReadToString(filePath)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(str), nil
}

//ReadToUint64 当文件内容为无符号单数字时，读取并格式化为uint64
func ReadToUint64(filePath string) (uint64, error) {
	content, err := ReadToTrimString(filePath)
	if err != nil {
		return 0, err
	}

	var ret uint64
	if ret, err = strconv.ParseUint(content, 10, 64); err != nil {
		return 0, err
	}
	return ret, nil
}

//ReadToInt64 当文件内容为有符号数字时，读取并格式化为int64
func ReadToInt64(filePath string) (int64, error) {
	content, err := ReadToTrimString(filePath)
	if err != nil {
		return 0, err
	}

	var ret int64
	if ret, err = strconv.ParseInt(content, 10, 64); err != nil {
		return 0, err
	}
	return ret, nil
}
