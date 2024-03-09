package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 1.一次性读取：适用于文件内容较少的情况
	// byteData, err := os.ReadFile("21-文件操作\\hello.txt") //使用的是相对路径
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(byteData))

	// 2.分片读取
	file, err := os.Open("21-文件操作\\hello.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// for {
	// 	var byteData = make([]byte, 13) //指定一次读取的容量，
	// 	l, err := file.Read(byteData)   //返回读取的字节数以及一个err
	// 	if err == io.EOF {              //表示文件读取完毕
	// 		break
	// 	}
	// 	fmt.Println(string(byteData[:l]))

	// }

	// 3.带缓冲读取
	// 3.1按行读取
	// buf := bufio.NewReader(file)
	// for {
	// 	line, _, err := buf.ReadLine()
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	fmt.Println(string(line), err)
	// }

	// 3.2指定分隔符读取
	buf := bufio.NewScanner(file)
	// 调用对应的方法指定分割符
	// buf.Split(bufio.ScanWords) //按单词分割
	// buf.Split(bufio.ScanRunes)//按Unicode字符分割
	// buf.Split(bufio.ScanLines)//按行分割
	// buf.Split(bufio.ScanBytes)//按字节分割
	// 可仿照提供的方法自定义分割函数
	buf.Split(ScanDefined)

	// 循环读取
	for buf.Scan() { //buf.Scan()返回一个bool值，判断文件是否已读取完
		fmt.Println(buf.Text()) //默认的读取方式是按行读取
	}
}

// 自定义分割函数
func ScanDefined(data []byte, atEOF bool) (advance int, token []byte, err error) {
	for i, v := range data {
		if v == '|' {
			return i + 1, data[:i], nil
		}
	}

	return
}
