package main

import (
	"fmt"
	"input/input"
	"io/ioutil"
)

func main() {
	dicts := []string{}
	// 查找词表文件夹下所有词表
	files, err := ioutil.ReadDir("./dicts")
	if err != nil {
		fmt.Println("read dict error")
		return
	}
	// 获取所有词表文件
	for _, file := range files {
		dicts = append(dicts, "./dicts/"+file.Name())
	}
	fmt.Println(dicts)
	// 读取所有词表
	instance := input.ReadDicts(dicts)

	input.Loop(instance)
}
