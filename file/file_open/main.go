package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

//利用切片读文件
func readFromFileBySlice() {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
	}

	//关闭文件
	defer fileObj.Close()
	//读文件
	// var tmp = make([]byte, 128) //指定读取长度
	// n, err := fileObj.Read(tmp)
	var tmp [128]byte
	for {
		n, err := fileObj.Read(tmp[:])
		if err != nil {
			fmt.Printf("read file failed, err:%v\n", err)
			return
		}
		fmt.Printf("读了%d个字节\n", n)
		// fmt.Println(string(tmp))
		fmt.Println(string(tmp[:n]))

		if n < 128 {
			return
		}
	}
}

//利用bufio读文件
func readFromFileByBufio() {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
	}

	//关闭文件
	defer fileObj.Close()

	reader := bufio.NewReader(fileObj)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Printf("read file failed, err:%v\n", err)
			return
		}
		fmt.Print(line)
	}
}

func readFromFileByIoutil() {
	ret, err := ioutil.ReadFile("./main.go")
	if err != nil {
		fmt.Printf(" failed, err:%v\n", err)
		return
	}
	fmt.Print(string(ret))
}

func main() {
	// readFromFileBySlice()
	// readFromFileByBufio()
	readFromFileByIoutil()
}
