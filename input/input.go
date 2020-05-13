package input

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"sort"
	"strconv"
	"strings"
)

// Node 汉字节点,成员根据排序要求设置
type Node struct {
	Word      string // 汉字
	Score     int64  // 汉字出现的频次
	Letter    string // 汉字对应字母
	WordIndex int64  // 汉字在词表中的位置
}

// Instance 输入法实例结构
type Instance struct {
	// Spell string
	Dicts map[string]map[string]*Node
}

// Nodes 检索到的汉字节点数据类型
type Nodes []Node

// MatchNodes 检索到的汉字节点
var MatchNodes = make(Nodes, 0)

// Len 计算待排序slice长度
func (nodes Nodes) Len() int {
	return len(nodes)
}

// Less 设置排序策略
// 不同频次的汉字，频次越高的排在越前面
// 相同频次的汉字，根据对应的拼音的字母序排列，字母序越小的排在越前面。
// 相同频次的汉字，对应的拼音字母序也相同，则根据文件中的顺序排列。
func (nodes Nodes) Less(i, j int) bool {
	if nodes[i].Score != nodes[j].Score {
		return nodes[i].Score > nodes[j].Score
	}
	if nodes[i].Letter != nodes[j].Letter {
		return nodes[i].Letter < nodes[j].Letter
	}
	if nodes[i].WordIndex != nodes[j].WordIndex {
		return nodes[i].WordIndex < nodes[j].WordIndex
	}
	return true
}

// Swap 交换位置
func (nodes Nodes) Swap(i, j int) {
	nodes[i], nodes[j] = nodes[j], nodes[i]
}

// ChanNode 用于协程间通信
var ChanNode = make(chan Nodes, 100)

// maxCount 允许返回的最大汉字个数
var maxCount = 10

// ReadDicts 读取词表
func ReadDicts(dicts []string) *Instance {
	instance := new(Instance)
	instance.Dicts = make(map[string]map[string]*Node)
	for _, dict := range dicts {
		var index int64 = 0
		file, err := os.Open(dict)
		if err != nil {
			fmt.Println("open " + dict + " fail")
			continue
		}
		buf := bufio.NewReader(file)
		for {
			filename := strings.Split(path.Base(dict), ".")[0]
			line, err := buf.ReadString('\n')
			if err != nil {
				// 当前词表读取完毕
				if err == io.EOF {
					fmt.Println("last line=" + line)
					break
				}
				fmt.Println("read dict error,err=" + err.Error())
				continue
			}
			line = strings.TrimSpace(line)
			// 如果是空行，跳过
			if len(line) == 0 {
				continue
			}

			item := strings.Split(line, " ")
			// 如果格式错误，跳过
			if len(item) != 2 {
				fmt.Println("the file format error")
				continue
			}
			if _, ok := instance.Dicts[filename]; !ok {
				instance.Dicts[filename] = make(map[string]*Node)
			}

			// fmt.Println(item)
			score, err := strconv.ParseInt(item[1], 10, 32)
			instance.Dicts[filename][item[0]] = new(Node)
			instance.Dicts[filename][item[0]].Word = item[0]
			instance.Dicts[filename][item[0]].Score = score
			instance.Dicts[filename][item[0]].WordIndex = index
			instance.Dicts[filename][item[0]].Letter = filename
			index++
		}
	}
	return instance
}

func findwords(dict map[string]*Node) {
	index := 10
	nodes := Nodes{}
	for _, val := range dict {
		nodes = append(nodes, *val)
	}
	sort.Sort(nodes)
	// lock.Lock()

	if len(nodes) < 10 {
		index = len(nodes)
	}
	// MatchNodes = append(MatchNodes, nodes[:index]...)
	ChanNode <- nodes[:index]
	// lock.Unlock()
	return
}

// FindWords 根据输入的单个汉字拼音，返回对应的汉字
func (instance *Instance) FindWords(spell string) (words []string) {
	cnt := 0
	have := false
	// 每次查找都要清空
	words = []string{}
	MatchNodes = Nodes{}
	// 处理输入的拼音
	spell = strings.TrimSpace(spell)
	if len(spell) == 0 {
		return words
	}

	// 如果输入的是完整拼音
	if mapNodes, ok := instance.Dicts[spell]; ok {
		go findwords(mapNodes)
		MatchNodes = <-ChanNode
		sort.Sort(MatchNodes)
		for _, val := range MatchNodes {
			words = append(words, val.Word)
		}
	} else {
		// 非完整拼音
		for py, word := range instance.Dicts {
			if strings.HasPrefix(py, spell) {
				cnt++
				go findwords(word)
				have = true
			}
		}
		// 如果检索的汉字，从channel中获取数据，并排序
		if have {
			for i := 0; i < cnt; i++ {
				MatchNodes = append(MatchNodes, <-ChanNode...)
			}
			sort.Sort(MatchNodes)
			for _, val := range MatchNodes {
				words = append(words, val.Word)
			}
		}
	}
	if len(words) >= maxCount {
		return words[:maxCount]
	}
	return words
}

// Loop 循环获取输入，并查找汉字
func Loop(instance *Instance) {
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
