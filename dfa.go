package goutil

import (
	"regexp"
	"strings"
)

// Node 结构
type Node struct {
	childrens map[string]*Node
	terminate bool
}

// Dfa DFA 算法
type Dfa struct {
	root Node
}

// 添加敏感词
// @word敏感词
func (p *Dfa) addWord(word string) {
	node := &p.root
	str := []rune(word)
	lastPos := len(str) - 1
	for i, s := range str {
		sKey := string(s)
		if node.childrens[sKey] == nil {
			node.childrens[sKey] = &Node{map[string]*Node{}, i == lastPos}
		} else {
			if i == lastPos {
				node.childrens[sKey] = &Node{map[string]*Node{}, true}
			}
		}
		node = node.childrens[sKey]
	}
}

// BuildTree 构建敏感词过滤器
// @words 敏感词列表
func (p *Dfa) BuildTree(words []string) {
	p.root = Node{map[string]*Node{}, false}
	for _, word := range words {
		p.addWord(word)
	}
}

// IsTreeBuild 判断过滤的树是否已经生成了
func (p *Dfa) IsTreeBuild() bool {
	var isBuild bool

	if len(p.root.childrens) > 0 {
		isBuild = true
	}

	return isBuild
}

// IsContain 判断是否包含敏感词
// @text 输入文本
func (p *Dfa) IsContain(text string) bool {
	str := []rune(text)
	strLen := len(str)
	for i := range str {
		j := i
		node := &p.root
		for j < strLen {
			currentWord := string(str[j])
			if node.childrens[currentWord] == nil {
				break
			} else {
				if node.childrens[currentWord].terminate {
					return true
				}
				node = node.childrens[currentWord]
			}
			j++
		}
	}
	return false
}

// FilterWords 过滤关键词
// @text 输入文本
// @replaceChar 替换的字符
func (p *Dfa) FilterWords(text, replaceChar string) (string, bool) {
	filterFlag := false
	str := []rune(text)
	strLen := len(str)
	strs := []string{}
	//是否继续添加
	isAppend := true
	for i := 0; i < strLen; i++ {
		node := &p.root
		for j := i; j < strLen && len(node.childrens) > 0; j++ {
			currentWord := string(str[j])
			if node.childrens[currentWord] == nil {
				break
			} else {
				if node.childrens[currentWord].terminate {
					isAppend = false
					temp := []string{}
					for k := 0; k < j-i+1; k++ {
						temp = append(temp, replaceChar)
					}
					strs = append(strs, strings.Join(temp, ""))
					i = j
					filterFlag = true
				}
				node = node.childrens[currentWord]
			}
		}
		if isAppend {
			strs = append(strs, string(str[i]))
		} else {
			isAppend = true
		}
	}
	result, flag := p.RegularFilter(strings.Join(strs, ""))
	return result, filterFlag || flag
}

// RegularFilter 正则过滤
func (p *Dfa) RegularFilter(text string) (string, bool) {
	newStr := text
	replaceFlag := false

	// 正则表达式
	regs := []*regexp.Regexp{
		regexp.MustCompile("0?(13[0-9]|15[012356789]|17[678]|18[0-9]|14[57])[0-9]{8}"),
		regexp.MustCompile("[1-9]{4,12}"),
	}
	for _, v := range regs {
		if v.MatchString(newStr) {
			newStr = v.ReplaceAllString(newStr, "****")
			replaceFlag = true
		}
	}
	return newStr, replaceFlag
}
