package test

import (
	"fmt"
	"input/input"
	"reflect"
	"testing"
)

// Node 汉字节点,成员根据排序要求设置
// type Node struct {
// 	Word      string // 汉字
// 	Score     int64  // 汉字出现的频次
// 	Letter    string // 汉字对应字母
// 	WordIndex int64  // 汉字在词表中的位置
// }
// type DictsMap map[string]map[string]*Node
var instance = new(input.Instance)
var dicts = map[string]map[string]*input.Node{
	"bai": {
		"白": &input.Node{
			"白",
			10,
			"bai",
			0,
		},
		"百": &input.Node{
			"百",
			9,
			"bai",
			1,
		},
		"摆": &input.Node{
			"摆",
			7,
			"bai",
			2,
		},
		"败": &input.Node{
			"败",
			4,
			"bai",
			3,
		},
		"柏": &input.Node{
			"柏",
			2,
			"bai",
			4,
		},
		"伯": &input.Node{
			"伯",
			1,
			"bai",
			5,
		},
		"拜": &input.Node{
			"拜",
			7,
			"bai",
			6,
		},
		"敗": &input.Node{
			"敗",
			1,
			"bai",
			7,
		},
		"掰": &input.Node{
			"掰",
			1,
			"bai",
			8,
		},
		"擺": &input.Node{
			"擺",
			1,
			"bai",
			9,
		},
		"拝": &input.Node{
			"拝",
			1,
			"bai",
			10,
		},
		"粨": &input.Node{
			"粨",
			1,
			"bai",
			11,
		},
	},
	"du": {
		"读": &input.Node{
			"读",
			10,
			"du",
			0,
		},
		"都": &input.Node{
			"都",
			9,
			"du",
			1,
		},
		"度": &input.Node{
			"度",
			8,
			"du",
			2,
		},
		"杜": &input.Node{
			"杜",
			7,
			"du",
			3,
		},
		"毒": &input.Node{
			"毒",
			5,
			"du",
			4,
		},
		"堵": &input.Node{
			"堵",
			6,
			"du",
			5,
		},
		"独": &input.Node{
			"独",
			10,
			"du",
			6,
		},
	},
	"ba": {
		"吧": &input.Node{
			"吧",
			7,
			"ba",
			0,
		},
		"把": &input.Node{
			"把",
			9,
			"ba",
			1,
		},
		"疤": &input.Node{
			"疤",
			10,
			"ba",
			2,
		},
		"爸": &input.Node{
			"爸",
			10,
			"ba",
			3,
		},
		"八": &input.Node{
			"八",
			7,
			"ba",
			4,
		},
		"拔": &input.Node{
			"拔",
			8,
			"ba",
			5,
		},
	},
	"cai": {
		"才": &input.Node{
			"才",
			10,
			"cai",
			0,
		},
		"菜": &input.Node{
			"菜",
			10,
			"cai",
			1,
		},
		"踩": &input.Node{
			"踩",
			5,
			"cai",
			2,
		},
		"猜": &input.Node{
			"猜",
			8,
			"cai",
			3,
		},
	},
	"cao": {
		"草": &input.Node{
			"草",
			7,
			"cao",
			0,
		},
		"操": &input.Node{
			"操",
			10,
			"cao",
			1,
		},
		"曹": &input.Node{
			"曹",
			6,
			"cao",
			2,
		},
	},
	"da": {
		"大": &input.Node{
			"大",
			7,
			"da",
			0,
		},
		"打": &input.Node{
			"打",
			10,
			"da",
			1,
		},
		"搭": &input.Node{
			"搭",
			6,
			"da",
			2,
		},
	},
	"zhan": {
		"展": &input.Node{
			"展",
			9,
			"zhan",
			0,
		},
		"战": &input.Node{
			"战",
			6,
			"zhan",
			1,
		},
		"站": &input.Node{
			"站",
			6,
			"zhan",
			2,
		},
	},
	"zhang": {
		"长": &input.Node{
			"长",
			10,
			"zhang",
			0,
		},
		"张": &input.Node{
			"张",
			9,
			"zhang",
			1,
		},
	},
}

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
	}
	for _, test := range tests {
		word := instance.FindWords(test.input)
		fmt.Println(test.input)
		if reflect.DeepEqual(word, test.output) {
			t.Errorf("dict.FindWords(%v) = %q", test.input, test.output)
		}
	}
}
