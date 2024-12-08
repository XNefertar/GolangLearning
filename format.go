// Golang
// pakage 包声明
// package main表示一个可独立执行的程序
// 每个 Go 应用程序都包含一个名为 main 的包
package main

// import 引入包
// fmt 格式化IO包 (format)
// C++20 中有类似头文件 <format>
import "fmt"

// func 函数
// main 主体函数
// 每个可执行程序所必须包含的
func main() {
	// Println 将字符串输出到控制台
	// 该函数会在结尾自动添加换行符
	fmt.Println("Golang")
	// 也可以使用Print函数
	// 但是需要显式的添加换行符\n
	fmt.Print("Hello World\n")
}
