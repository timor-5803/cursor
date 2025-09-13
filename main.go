package main

import "fmt"

func main() {
	// 在这里设置断点测试
	fmt.Println("Hello, World!")

	// 添加一些变量用于调试
	name := "Go Debugger"
	version := "1.0"

	fmt.Printf("Welcome to %s version %s\n", name, version)

	// 在这里也可以设置断点
	fmt.Println("Debug session started!")
}
