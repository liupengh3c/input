package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"input/input"
)

// main the function where execution of the program begins
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

	// 读取所有词表
	instance := input.ReadDicts(dicts)

	Loop(instance)
}

// Loop 循环获取输入，并查找汉字
func Loop(instance *input.Instance) {
	stdin := bufio.NewReader(os.Stdin)
	for {
		spell, err := stdin.ReadString('\n')
		if err != nil {
			break
		}
		spell = strings.TrimRight(spell, "\n")
		spell = strings.TrimSpace(spell)
		words := instance.FindWords(spell)
		fmt.Println(strings.Join(words, ", "))
	}
}
