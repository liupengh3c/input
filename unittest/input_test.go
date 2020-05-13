package unittest

import (
	"fmt"
	"input/input"
	"reflect"
	"testing"
)

var instance = new(input.Instance)
var dicts = map[string]map[string]*input.Node{
	"bai": {
		"白": &input.Node{
			Word:      "白",
			Score:     10,
			Letter:    "bai",
			WordIndex: 0,
		},
		"百": &input.Node{
			Word:      "百",
			Score:     9,
			Letter:    "bai",
			WordIndex: 1,
		},
		"摆": &input.Node{
			Word:      "摆",
			Score:     7,
			Letter:    "bai",
			WordIndex: 2,
		},
		"败": &input.Node{
			Word:      "败",
			Score:     4,
			Letter:    "bai",
			WordIndex: 3,
		},
		"柏": &input.Node{
			Word:      "柏",
			Score:     2,
			Letter:    "bai",
			WordIndex: 4,
		},
		"伯": &input.Node{
			Word:      "伯",
			Score:     1,
			Letter:    "bai",
			WordIndex: 5,
		},
		"拜": &input.Node{
			Word:      "拜",
			Score:     7,
			Letter:    "bai",
			WordIndex: 6,
		},
		"敗": &input.Node{
			Word:      "敗",
			Score:     1,
			Letter:    "bai",
			WordIndex: 7,
		},
		"掰": &input.Node{
			Word:      "掰",
			Score:     1,
			Letter:    "bai",
			WordIndex: 8,
		},
		"擺": &input.Node{
			Word:      "擺",
			Score:     1,
			Letter:    "bai",
			WordIndex: 9,
		},
		"拝": &input.Node{
			Word:      "拝",
			Score:     1,
			Letter:    "bai",
			WordIndex: 10,
		},
		"粨": &input.Node{
			Word:      "粨",
			Score:     1,
			Letter:    "bai",
			WordIndex: 11,
		},
	},
	"du": {
		"读": &input.Node{
			Word:      "读",
			Score:     10,
			Letter:    "du",
			WordIndex: 0,
		},
		"都": &input.Node{
			Word:      "都",
			Score:     9,
			Letter:    "du",
			WordIndex: 1,
		},
		"度": &input.Node{
			Word:      "度",
			Score:     8,
			Letter:    "du",
			WordIndex: 2,
		},
		"杜": &input.Node{
			Word:      "杜",
			Score:     7,
			Letter:    "du",
			WordIndex: 3,
		},
		"毒": &input.Node{
			Word:      "毒",
			Score:     5,
			Letter:    "du",
			WordIndex: 4,
		},
		"堵": &input.Node{
			Word:      "堵",
			Score:     6,
			Letter:    "du",
			WordIndex: 5,
		},
		"独": &input.Node{
			Word:      "独",
			Score:     10,
			Letter:    "du",
			WordIndex: 6,
		},
	},
	"ba": {
		"吧": &input.Node{
			Word:      "吧",
			Score:     7,
			Letter:    "ba",
			WordIndex: 0,
		},
		"把": &input.Node{
			Word:      "把",
			Score:     9,
			Letter:    "ba",
			WordIndex: 1,
		},
		"疤": &input.Node{
			Word:      "疤",
			Score:     10,
			Letter:    "ba",
			WordIndex: 2,
		},
		"爸": &input.Node{
			Word:      "爸",
			Score:     10,
			Letter:    "ba",
			WordIndex: 3,
		},
		"八": &input.Node{
			Word:      "八",
			Score:     7,
			Letter:    "ba",
			WordIndex: 4,
		},
		"拔": &input.Node{
			Word:      "拔",
			Score:     8,
			Letter:    "ba",
			WordIndex: 5,
		},
	},
	"cai": {
		"才": &input.Node{
			Word:      "才",
			Score:     10,
			Letter:    "cai",
			WordIndex: 0,
		},
		"菜": &input.Node{
			Word:      "菜",
			Score:     10,
			Letter:    "cai",
			WordIndex: 1,
		},
		"踩": &input.Node{
			Word:      "踩",
			Score:     5,
			Letter:    "cai",
			WordIndex: 2,
		},
		"猜": &input.Node{
			Word:      "猜",
			Score:     8,
			Letter:    "cai",
			WordIndex: 3,
		},
	},
	"cao": {
		"草": &input.Node{
			Word:      "草",
			Score:     7,
			Letter:    "cao",
			WordIndex: 0,
		},
		"操": &input.Node{
			Word:      "操",
			Score:     10,
			Letter:    "cao",
			WordIndex: 1,
		},
		"曹": &input.Node{
			Word:      "曹",
			Score:     6,
			Letter:    "cao",
			WordIndex: 2,
		},
	},
	"da": {
		"大": &input.Node{
			Word:      "大",
			Score:     7,
			Letter:    "da",
			WordIndex: 0,
		},
		"打": &input.Node{
			Word:      "打",
			Score:     10,
			Letter:    "da",
			WordIndex: 1,
		},
		"搭": &input.Node{
			Word:      "搭",
			Score:     6,
			Letter:    "da",
			WordIndex: 2,
		},
	},
	"zhan": {
		"展": &input.Node{
			Word:      "展",
			Score:     9,
			Letter:    "zhan",
			WordIndex: 0,
		},
		"战": &input.Node{
			Word:      "战",
			Score:     6,
			Letter:    "zhan",
			WordIndex: 1,
		},
		"站": &input.Node{
			Word:      "站",
			Score:     6,
			Letter:    "zhan",
			WordIndex: 2,
		},
	},
	"zhang": {
		"长": &input.Node{
			Word:      "长",
			Score:     10,
			Letter:    "zhang",
			WordIndex: 0,
		},
		"张": &input.Node{
			Word:      "张",
			Score:     9,
			Letter:    "zhang",
			WordIndex: 1,
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
		fmt.Printf("dict.FindWords(%v) = %q,expect output=%q\n", test.input, word, test.output)
		if !reflect.DeepEqual(word, test.output) {
			t.Errorf("dict.FindWords(%v) = %q,expect output=%q", test.input, word, test.output)
		}
	}
}
