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
	// 注意，不同于C/C++
	// 该语言行尾不需要加上;
	// 该行为由golang编译器自主完成

	// Println 将字符串输出到控制台
	// 该函数会在结尾自动添加换行符
	fmt.Println("Golang")
	// 也可以使用Print函数
	// 但是需要显式的添加换行符\n
	fmt.Print("Hello World\n")

	// 格式化字符串(可以参考C/C++的sprintf 和 printf)
	// Sprintf
	var stockcode1 = 123
	var enddate1 = "2020-12-31"
	var url1 = "Code=%d&endDate=%s"
	var target_url = fmt.Sprintf(url1, stockcode1, enddate1)
	fmt.Println(target_url)

	// Printf
	var stockcode2 = 123
	var enddate2 = "2020-12-31"
	var url2 = "Code=%d&endDate=%s"
	fmt.Printf(url2, stockcode2, enddate2)

}
