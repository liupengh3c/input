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

// ChanNode 用于协程间通信
// var ChanNode = make(chan Nodes)

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
	} else if nodes[i].Letter != nodes[j].Letter {
		return nodes[i].Letter < nodes[j].Letter
	} else {
		return nodes[i].WordIndex < nodes[j].WordIndex
	}
}

// Swap 交换位置
func (nodes Nodes) Swap(i, j int) {
	nodes[i], nodes[j] = nodes[j], nodes[i]
}

// maxCount 允许返回的最大汉字个数
var maxCount = 10

// ReadDictFile 读取单个词表文件，并返回所有词表
func ReadDictFile(name string) ([]string, error) {
	dicts := make([]string, 0)
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	buf := bufio.NewReader(file)
	for {
		line, err := buf.ReadString('\n')
		if err != nil {
			// 当前词表读取完毕
			if err == io.EOF {
				fmt.Println("last line=" + line)
				break
			}
		}
		line = strings.TrimSpace(line)
		// 如果是空行，跳过
		if len(line) == 0 {
			continue
		}
		item := strings.Split(line, " ")
		// 如果格式错误，跳过
		if len(item) != 2 {
			// fmt.Println("the file format error,file=" + name)
			continue
		}
		dicts = append(dicts, line)
	}
	return dicts, nil
}

// ReadDicts 读取词表
func ReadDicts(dicts []string) *Instance {
	instance := new(Instance)
	instance.Dicts = make(map[string]map[string]*Node)
	for _, dict := range dicts {
		var index int64 = 0
		sliDict, err := ReadDictFile(dict)
		if err != nil {
			continue
		}
		// 获取到词典对应的拼音
		pinyin := strings.Split(path.Base(dict), ".")[0]
		if _, ok := instance.Dicts[pinyin]; !ok {
			instance.Dicts[pinyin] = make(map[string]*Node)
		}
		for _, line := range sliDict {
			item := strings.Split(line, " ")
			// 汉字
			word := item[0]
			// 频率分值
			score, err := strconv.ParseInt(item[1], 10, 32)
			if err != nil {
				continue
			}
			node := &Node{
				Word:      item[0],
				Score:     score,
				WordIndex: index,
				Letter:    pinyin,
			}
			instance.Dicts[pinyin][word] = node
			index++
		}
	}
	return instance
}

// Findword 在某个拼音下查找汉字
func Findword(dict map[string]*Node, ch chan<- Nodes) {
	index := maxCount
	nodes := Nodes{}
	for _, val := range dict {
		nodes = append(nodes, *val)
	}
	sort.Sort(nodes)

	if len(nodes) < maxCount {
		index = len(nodes)
	}

	ch <- nodes[:index]
}

// FindWords 根据输入的单个汉字拼音，返回对应的汉字
func (instance *Instance) FindWords(spell string) (words []string) {
	words = []string{}
	// 处理输入的拼音,spell为空代表没有输入，直接返回
	spell = strings.TrimSpace(spell)
	if len(spell) == 0 {
		return words
	}

	cnt := 0
	words = []string{}
	MatchNodes = Nodes{}
	ch := make(chan Nodes)

	// 如果输入的是完整拼音
	if mapNodes, ok := instance.Dicts[spell]; ok {
		go Findword(mapNodes, ch)
		// MatchNodes = <-ChanNode
		MatchNodes = <-ch
		sort.Sort(MatchNodes)
		for _, val := range MatchNodes {
			words = append(words, val.Word)
		}
	} else {
		// 非完整拼音,在所有前缀与输入相同的拼音的汉字中并发去query
		for py, word := range instance.Dicts {
			if strings.HasPrefix(py, spell) {
				cnt++
				go Findword(word, ch)
			}
		}
		// 如果检索的汉字，从channel中获取数据，并排序
		if cnt > 0 {
			for i := 0; i < cnt; i++ {
				MatchNodes = append(MatchNodes, <-ch...)
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
