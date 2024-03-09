package main

import (
	"fmt"
	"os"
)

// O_RDONLY int = syscall.O_RDONLY // open the file read-only.
// O_WRONLY int = syscall.O_WRONLY // open the file write-only.
// O_RDWR   int = syscall.O_RDWR   // open the file read-write.
// // The remaining values may be or'ed in to control behavior.
// O_APPEND int = syscall.O_APPEND // append data to the file when writing.
// O_CREATE int = syscall.O_CREAT  // create a new file if none exists.
// O_EXCL   int = syscall.O_EXCL   // used with O_CREATE, file must not exist.
// O_SYNC   int = syscall.O_SYNC   // open for synchronous I/O.
// O_TRUNC  int = syscall.O_TRUNC  // truncate regular writable file when opened.

func main() {
	// // 1.文件写入：第二个参数指定打开文件的模式（可通过 | 来同时赋予多个权限），第三个参数设置文件的权限，在Windows中没有实际意义
	// file, err := os.OpenFile("21-文件操作\\write.txt", os.O_CREATE|os.O_RDWR, 0721)
	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()

	// file.Write([]byte("这是写入的内容~"))

	// 2.一次性写入：写入的内容会覆盖该文件已有的内容
	// err := os.WriteFile("21-文件操作\\write.txt", []byte("这是一次性写入的内容"), 0721)
	// fmt.Println(err)

	// // 3.文件拷贝
	// rfile, err := os.Open("C:\\Users\\31715\\Desktop\\new_file\\bg\\lqx.jpg")
	// if err != nil {
	// 	panic(err)
	// }

	// defer rfile.Close()

	// wfile, err := os.OpenFile("21-文件操作\\lqx.jpg", os.O_CREATE|os.O_RDWR, 0721)
	// if err != nil {
	// 	panic(err)
	// }
	// defer wfile.Close()

	// // 第一个参数是拷贝生成的文件，第二个参数是被拷贝的源文件
	// io.Copy(wfile, rfile)

	// 3.文件目录处理
	// 获取对应的文件目录
	dir, err := os.ReadDir("21-文件操作\\")
	if err != nil {
		panic(err)
	}

	// 遍历目录下的每个文件
	for _, v := range dir {
		// 调用对应方法获取文件的信息
		info, _ := v.Info()
		fmt.Println(v.Name(), v.IsDir(), info.Size())
	}
}
