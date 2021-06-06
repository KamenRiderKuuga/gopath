package main

import "fmt"

func printMap(cityMap map[string]string) {
	// 引用传递
	for key, value := range cityMap {
		fmt.Println("key =", key, "value =", value)
	}
}

func changeValue(cityMap map[string]string) {
	cityMap["England"] = "London"
}

func main() {
	cityMap := make(map[string]string)

	// 添加key,value
	cityMap["China"] = "Beijing"
	cityMap["Japan"] = "Tokyo"
	cityMap["USA"] = "NewYork"

	printMap(cityMap)

	// 删除
	delete(cityMap, "China")

	// 修改
	cityMap["USA"] = "DC"
	changeValue(cityMap)

	fmt.Println("----------")

	printMap(cityMap)
}
