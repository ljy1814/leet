package main

import "testing"

func Benchmark140(b *testing.B) {
	s := "catsanddog"
	wordDict := []string{"cat", "cats", "and", "sand", "dog"}
	for i := 0; i < b.N; i++ {
		wordBreak(s, wordDict)
	}
}

func Benchmark141(b *testing.B) {
	s := "catsanddog"
	wordDict := []string{"cat", "cats", "and", "sand", "dog"}
	for i := 0; i < b.N; i++ {
		wordBreak1(s, wordDict)
	}
}

func Benchmark142(b *testing.B) {
	s := "catsanddog"
	wordDict := []string{"cat", "cats", "and", "sand", "dog"}
	//s = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	//wordDict = []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}
	for i := 0; i < b.N; i++ {
		wordBreak0(s, wordDict)
	}
}
