package main

import "fmt"

const (
	SuccessCode    Code = 0
	ServiceErrCode Code = 2001
	NetWorkErrCode Code = 3001
)

// 定义一个自定义类型,底层是int
type Code int

// 为自定义类型绑定一个方法
func (c Code) getMsg() string {
	switch c {
	case NetWorkErrCode:
		return "网络错误"
	case ServiceErrCode:
		return "服务错误"
	}
	return "成功"
}

func (c Code) getResult() (code Code, msg string) {
	return c, c.getMsg()
}

func WebServer(name int) (code Code, msg string) {

	if name == 2 {
		return NetWorkErrCode.getResult()
	}
	if name == 3 {
		return ServiceErrCode.getResult()
	}

	return SuccessCode.getResult()

}

func main() {
	code := 3
	fmt.Println(WebServer(code))
}
