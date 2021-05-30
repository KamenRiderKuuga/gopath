package main

import (
	"GolangStudy/5-init/lib1"

	// 匿名导包，可以在当前包不引用导入包的函数，依然会走被导入包的init函数
	// _ "GolangStudy/5-init/lib1"

	// 指定包名称进行导包，使用导入包函数的时候使用指定的包名称来调用函数
	// l1 "GolangStudy/5-init/lib1"

	// 使用这种方式导包，在当前包内调用被导入包函数时可以不指定被导入包的包名
	// . "GolangStudy/5-init/lib1"
	"GolangStudy/5-init/lib2"
)

func main() {
	lib1.Lib1Test()
	lib2.Lib2Test()
}
