package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func writeFile1() {
	fileObj, err := os.OpenFile("./xx.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}
	defer fileObj.Close()

	//write
	fileObj.Write([]byte("写入切片\n"))
	fileObj.WriteString("写入字符串\n")
}

func writeFileByBufio() {
	fileObj, err := os.OpenFile("./xx.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}
	defer fileObj.Close()

	//write
	writer := bufio.NewWriter(fileObj)
	for i := 0; i < 5; i++ {
		writer.WriteString("bufio写入\n")
	}
	writer.Flush()
}

func writeFileByIoutil() {
	str := "ioutil写入文件"
	err := ioutil.WriteFile("./xx.txt", []byte(str), 0644)
	if err != nil {
		fmt.Printf("write file failed, err:%v\n", err)
		return
	}
}

func CopyFile(dstName string, srcName string) (written int64, err error) {
	//只读方式打开文件
	src, err := os.Open(srcName)
	if err != nil {
		fmt.Printf("open %s failed, err:%v\n", srcName, err)
		return
	}
	defer src.Close()
	//以写|创建模式打开目标文件
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("open %s failed, err:%v\n", dstName, err)
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}

func main() {
	// writeFile1()
	// writeFileByBufio()
	// writeFileByIoutil()
	CopyFile("./copy.txt", "./xx.txt")
}
