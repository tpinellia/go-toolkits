package file

import (
	"fmt"
	"testing"
)

func TestPath(t *testing.T) {
	fmt.Println(GetSelfAbsolutePath())
	fmt.Println(GetSelfAbsoluteDir())
	fmt.Println(GetFilePath("../logger/logger.go"))
	fmt.Println(GetFilePath("/tmp1"))
	fmt.Println(GetFileDir("logger.go"))
	fmt.Println(GetCurrentDir())
}

func TestFile(t *testing.T) {
	fmt.Println(IsFile("./file.go"))
	fmt.Println(IsExist("/tmp1"))
}

func TestRead(t *testing.T) {
	fmt.Println(ReadToString("./write.go"))
}
