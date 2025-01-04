package main

import (
	"fmt"
	"testing"
)

// You are given a string num consisting of only digits. A string of digits is called balanced if the sum of the digits at even indices is equal to the sum of digits at odd indices.
// Return true if num is balanced, otherwise return false.
// Example 1:
// Input: num = "1234"
// Output: false

// Explanation:
// The sum of digits at even indices is 1 + 3 == 4, and the sum of digits at odd indices is 2 + 4 == 6.
// Since 4 is not equal to 6, num is not balanced.
// Example 2:
// Input: num = "24123"
// Output: true
// Explanation:
// The sum of digits at even indices is 2 + 1 + 3 == 6, and the sum of digits at odd indices is 4 + 2 == 6.
// Since both are equal the num is balanced.

func TestBalancedDigit(t *testing.T) {
	num := "1234"
	balanced_digit(num)
	num = "24123"
	balanced_digit(num)
}

func TestReverseWord(t *testing.T) {
	sent := "hello world welcome"
	reverse_word(sent)
}

// Find All Anagrams in a String
// Medium
// Topics
// Companies
// Given two strings s and p, return an array of all the start indices of p's
// anagrams
//  in s. You may return the answer in any order.

// Example 1:
// Input: s = "cbaebabacd", p = "abc"
// Output: [0,6]
// Explanation:
// The substring with start index = 0 is "cba", which is an anagram of "abc".
// The substring with start index = 6 is "bac", which is an anagram of "abc".

// func findAnagram(s string , p string)
func TestFindAnagram(t *testing.T) {
	s := "cbaebabacdacb"
	p := "abc"
	// s := "abab"
	// p := "ab"
	res := findAnagram(s, p)
	fmt.Println(res)
}

/*
Example 1:Given an array of strings strs, group the
anagrams together. You can return the answer in any order.

Input: strs = ["eat","tea","tan","ate","nat","bat"]

Output: [["bat"],["nat","tan"],["ate","eat","tea"]]

Explanation:
There is no string in strs that can be rearranged to form "bat".
The strings "nat" and "tan" are anagrams as they can be rearranged to form each other.
The strings "ate", "eat", and "tea" are anagrams as they can be rearranged to form each other.
*/

func TestGroupAnagrams(t *testing.T) {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	res := groupAnagrams2(strs)
	fmt.Println(res)
}
