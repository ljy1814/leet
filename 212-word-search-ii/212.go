package main

import "fmt"

/*
给定一个二维网格 board 和一个字典中的单词列表 words，找出所有同时在二维网格和字典中出现的单词。

单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母在一个单词中不允许被重复使用。

示例:

输入:
words = ["oath","pea","eat","rain"] and board =
[
  ['o','a','a','n'],
  ['e','t','a','e'],
  ['i','h','k','r'],
  ['i','f','l','v']
]

输出: ["eat","oath"]
*/

type TrieNode struct {
	word string
	ch   byte
	next map[byte]*TrieNode
}

func initTree(words []string) *TrieNode {
	root := &TrieNode{}
	for _, w := range words {
		n := root
		for _, c := range w {
			if n.next == nil {
				n.next = make(map[byte]*TrieNode)
			}
			b := byte(c)
			n.next[b] = &TrieNode{}
			n.ch = b
			n = n.next[b]
		}
		n.word = w
	}
	return root
}

func dfs(board [][]byte, i, j int, root *TrieNode, wordMap map[string]bool) {
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
		return
	}
	c := board[i][j]
	cur := root.next[c]
	if c == '#' || cur == nil {
		return
	}

	if cur.word != "" {
		wordMap[cur.word] = true
	}

	board[i][j] = '#'
	dfs(board, i-1, j, cur, wordMap)
	dfs(board, i+1, j, cur, wordMap)
	dfs(board, i, j-1, cur, wordMap)
	dfs(board, i, j+1, cur, wordMap)
	board[i][j] = c

	return
}

func findWords(board [][]byte, words []string) []string {
	root := initTree(words)
	wordMap := make(map[string]bool)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			dfs(board, i, j, root, wordMap)
			if len(wordMap) > len(words) {
				return words
			}
		}
	}

	res := make([]string, 0, len(wordMap))
	for w := range wordMap {
		res = append(res, w)
	}

	return res
}

func main() {
	words := []string{"oath", "pea", "eat", "rain"}
	board :=
		[][]byte{
			[]byte{'o', 'a', 'a', 'n'},
			[]byte{'e', 't', 'a', 'e'},
			[]byte{'i', 'h', 'k', 'r'},
			[]byte{'i', 'f', 'l', 'v'},
		}
	res := findWords(board, words)
	fmt.Println(res)
}
