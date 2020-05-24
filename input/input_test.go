package input

import (
	"path"
	"reflect"
	"sort"
	"strings"
	"testing"
)

var instance = new(Instance)

// 构建测试词典
var dicts = map[string]map[string]*Node{
	"bai": {
		"白": &Node{
			Word:      "白",
			Score:     10,
			Letter:    "bai",
			WordIndex: 0,
		},
		"百": &Node{
			Word:      "百",
			Score:     9,
			Letter:    "bai",
			WordIndex: 1,
		},
		"摆": &Node{
			Word:      "摆",
			Score:     7,
			Letter:    "bai",
			WordIndex: 2,
		},
		"败": &Node{
			Word:      "败",
			Score:     4,
			Letter:    "bai",
			WordIndex: 3,
		},
		"柏": &Node{
			Word:      "柏",
			Score:     2,
			Letter:    "bai",
			WordIndex: 4,
		},
		"伯": &Node{
			Word:      "伯",
			Score:     1,
			Letter:    "bai",
			WordIndex: 5,
		},
		"拜": &Node{
			Word:      "拜",
			Score:     7,
			Letter:    "bai",
			WordIndex: 6,
		},
		"敗": &Node{
			Word:      "敗",
			Score:     1,
			Letter:    "bai",
			WordIndex: 7,
		},
		"掰": &Node{
			Word:      "掰",
			Score:     1,
			Letter:    "bai",
			WordIndex: 8,
		},
		"擺": &Node{
			Word:      "擺",
			Score:     1,
			Letter:    "bai",
			WordIndex: 9,
		},
		"拝": &Node{
			Word:      "拝",
			Score:     1,
			Letter:    "bai",
			WordIndex: 10,
		},
		"粨": &Node{
			Word:      "粨",
			Score:     1,
			Letter:    "bai",
			WordIndex: 11,
		},
	},
	"du": {
		"读": &Node{
			Word:      "读",
			Score:     10,
			Letter:    "du",
			WordIndex: 0,
		},
		"都": &Node{
			Word:      "都",
			Score:     9,
			Letter:    "du",
			WordIndex: 1,
		},
		"度": &Node{
			Word:      "度",
			Score:     8,
			Letter:    "du",
			WordIndex: 2,
		},
		"杜": &Node{
			Word:      "杜",
			Score:     7,
			Letter:    "du",
			WordIndex: 3,
		},
		"毒": &Node{
			Word:      "毒",
			Score:     5,
			Letter:    "du",
			WordIndex: 4,
		},
		"堵": &Node{
			Word:      "堵",
			Score:     6,
			Letter:    "du",
			WordIndex: 5,
		},
		"独": &Node{
			Word:      "独",
			Score:     10,
			Letter:    "du",
			WordIndex: 6,
		},
	},
	"ba": {
		"吧": &Node{
			Word:      "吧",
			Score:     7,
			Letter:    "ba",
			WordIndex: 0,
		},
		"把": &Node{
			Word:      "把",
			Score:     9,
			Letter:    "ba",
			WordIndex: 1,
		},
		"疤": &Node{
			Word:      "疤",
			Score:     10,
			Letter:    "ba",
			WordIndex: 2,
		},
		"爸": &Node{
			Word:      "爸",
			Score:     10,
			Letter:    "ba",
			WordIndex: 3,
		},
		"八": &Node{
			Word:      "八",
			Score:     7,
			Letter:    "ba",
			WordIndex: 4,
		},
		"拔": &Node{
			Word:      "拔",
			Score:     8,
			Letter:    "ba",
			WordIndex: 5,
		},
	},
	"cai": {
		"才": &Node{
			Word:      "才",
			Score:     10,
			Letter:    "cai",
			WordIndex: 0,
		},
		"菜": &Node{
			Word:      "菜",
			Score:     10,
			Letter:    "cai",
			WordIndex: 1,
		},
		"踩": &Node{
			Word:      "踩",
			Score:     5,
			Letter:    "cai",
			WordIndex: 2,
		},
		"猜": &Node{
			Word:      "猜",
			Score:     8,
			Letter:    "cai",
			WordIndex: 3,
		},
	},
	"cao": {
		"草": &Node{
			Word:      "草",
			Score:     7,
			Letter:    "cao",
			WordIndex: 0,
		},
		"操": &Node{
			Word:      "操",
			Score:     10,
			Letter:    "cao",
			WordIndex: 1,
		},
		"曹": &Node{
			Word:      "曹",
			Score:     6,
			Letter:    "cao",
			WordIndex: 2,
		},
	},
	"da": {
		"大": &Node{
			Word:      "大",
			Score:     7,
			Letter:    "da",
			WordIndex: 0,
		},
		"打": &Node{
			Word:      "打",
			Score:     10,
			Letter:    "da",
			WordIndex: 1,
		},
		"搭": &Node{
			Word:      "搭",
			Score:     6,
			Letter:    "da",
			WordIndex: 2,
		},
	},
	"zhan": {
		"展": &Node{
			Word:      "展",
			Score:     9,
			Letter:    "zhan",
			WordIndex: 0,
		},
		"战": &Node{
			Word:      "战",
			Score:     6,
			Letter:    "zhan",
			WordIndex: 1,
		},
		"站": &Node{
			Word:      "站",
			Score:     6,
			Letter:    "zhan",
			WordIndex: 2,
		},
	},
	"zhang": {
		"长": &Node{
			Word:      "长",
			Score:     10,
			Letter:    "zhang",
			WordIndex: 0,
		},
		"张": &Node{
			Word:      "张",
			Score:     9,
			Letter:    "zhang",
			WordIndex: 1,
		},
	},
	"fan": {
		"饭": &Node{
			Word:      "饭",
			Score:     10,
			Letter:    "fan",
			WordIndex: 0,
		},
		"烦": &Node{
			Word:      "烦",
			Score:     9,
			Letter:    "fan",
			WordIndex: 1,
		},
	},
	"fang": {
		"放": &Node{
			Word:      "放",
			Score:     10,
			Letter:    "fang",
			WordIndex: 0,
		},
		"方": &Node{
			Word:      "方",
			Score:     9,
			Letter:    "fang",
			WordIndex: 1,
		},
	},
}

// TestReadDicts 读取词典函数测试
func TestReadDicts(t *testing.T) {
	type words struct {
		name  string
		score int64
	}
	var test = struct {
		input  string
		output []words
	}{
		input: "./testdata/zhang.dat",
		output: []words{
			{
				name:  "长",
				score: 10,
			},
			{
				name:  "张",
				score: 9,
			},
		},
	}
	filename := path.Base(test.input)
	pinyin := strings.Split(filename, ".")[0]
	// 加入一个xxx错误路径，测试读入词典文件错误分支
	in := ReadDicts([]string{test.input, "xxx"})
	if len(test.output) != len(in.Dicts[pinyin]) {
		t.Errorf("length of output(%d) is not equal espect output length(%d)", len(in.Dicts[pinyin]), len(test.output))
		return
	}
	for _, word := range test.output {
		if node, ok := in.Dicts[pinyin][word.name]; !ok {
			t.Errorf("the espect word(%v) is not in dicts", word.name)
			break
		} else {
			if word.score != node.Score {
				t.Errorf("the espect word(%v)-score(%v) is not equal the dict score(%v)", word.name, word.score, node.Score)
				break
			}
		}
	}
}

// TestFindWords 检索汉字函数测试
func TestFindWords(t *testing.T) {
	// 构造测试词典
	instance.Dicts = dicts

	// test tables
	var tests = []struct {
		input  string
		output []string
	}{
		{
			"bai",
			[]string{"白", "百", "摆", "拜", "败", "柏", "伯", "敗", "掰", "擺"},
		},
		{
			"da",
			[]string{"打", "大", "搭"},
		},
		{
			"zha",
			[]string{"长", "展", "张", "战", "站"},
		},
		{
			"d",
			[]string{"打", "读", "独", "都", "度", "大", "杜", "搭", "堵", "毒"},
		},
		{
			"",
			[]string{},
		},
	}
	for _, test := range tests {
		word := instance.FindWords(test.input)
		// fmt.Printf("dict.FindWords(%v) = %q,expect output=%q\n", test.input, word, test.output)
		if !reflect.DeepEqual(word, test.output) {
			t.Errorf("dict.FindWords(%v) = %q,expect output=%q", test.input, word, test.output)
		}
	}
}

// TestFindWord 测试在某个拼音下查找汉字
func TestFindWord(t *testing.T) {
	words := make([]string, 0)
	var test = struct {
		input  string
		output []string
	}{
		"da",
		[]string{"打", "大", "搭"},
	}
	ch := make(chan Nodes)
	go Findword(dicts[test.input], ch)
	nodes := <-ch
	sort.Sort(nodes)
	for _, val := range nodes {
		words = append(words, val.Word)
	}
	if !reflect.DeepEqual(words, test.output) {
		t.Errorf("FindWord(%v) = %q,expect output=%q", test.input, words, test.output)
	}
}

// TestSort 测试排序功能
func TestSort(t *testing.T) {
	var nodes = Nodes{}
	words := make([]string, 0)
	var test = struct {
		input  string
		output []string
	}{
		"bai",
		[]string{"白", "百", "摆", "拜", "败", "柏", "伯", "敗", "掰", "擺", "拝", "粨"},
	}
	for _, word := range dicts[test.input] {
		nodes = append(nodes, *word)
	}
	sort.Sort(nodes)
	for _, val := range nodes {
		words = append(words, val.Word)
	}
	if !reflect.DeepEqual(words, test.output) {
		t.Errorf("sort output(%v) = %q,expect output=%q", test.input, words, test.output)
	}
}

func TestReadDictFile(t *testing.T) {
	var test = struct {
		input  string
		output []string
	}{
		input: "./testdata/zhang.dat",
		output: []string{
			"长 10",
			"张 9",
			"章 abc",
		},
	}
	res, err := ReadDictFile(test.input)
	if err != nil {
		t.Errorf("read file " + test.input + " fail,err=" + err.Error())
	}
	if !reflect.DeepEqual(res, test.output) {
		t.Errorf("read file(%v) = %q,expect output=%q", test.input, res, test.output)
	}
	test = struct {
		input  string
		output []string
	}{
		input: "./testdata/zhang1212.dat",
		output: []string{
			"长 10",
			"张 9",
			"章 abc",
		},
	}

	res, err = ReadDictFile(test.input)
	if err == nil {
		t.Errorf("read file " + test.input + " return success,espect fail")
	}
	test = struct {
		input  string
		output []string
	}{
		input: "./testdata/bai.dat",
		output: []string{
			"白 6",
			"百 9",
		},
	}

	res, err = ReadDictFile(test.input)
	if err != nil {
		t.Errorf("read file " + test.input + " fail,err=" + err.Error())
	}
	if !reflect.DeepEqual(res, test.output) {
		t.Errorf("read file(%v) = %q,expect output=%q", test.input, res, test.output)
	}
}
