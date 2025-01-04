package main

import (
	"fmt"
	"strings"
)

func balanced_digit(num string) bool {
	var sum1, sum2 int
	for i := 0; i < len(num); i++ {
		if i%2 == 0 {
			val := int(num[i] - '0')
			sum1 = sum1 + val
		} else {
			val := int(num[i] - '0')
			sum2 = sum2 + val
		}
	}
	if sum1 == sum2 {
		fmt.Printf("Balanced digit %s\n", num)
		return true
	} else {
		fmt.Printf("Not Balanced digit %s\n", num)
		return false
	}
}

func reverse_word(word string) {
	words := strings.Split(word, " ")
	j := len(words) - 1
	for i := 0; i < j; i++ {
		words[i], words[j] = words[j], words[i]
		j--
	}
	fmt.Println(words)
}

func findAnagram(s string, p string) []string {
	result := make([]string, 0)
	example := [26]int{}
	for _, word := range p {
		example[word-'a']++
	}
	fmt.Println(example)

	for i := 0; i <= len(s)-len(p); i++ {
		res := [26]int{}
		for _, word := range s[i:(i + len(p))] {
			res[word-'a']++
		}
		if res == example {
			result = append(result, s[i:i+len(p)])
		}
	}
	return result
}

func groupAnagrams(strs []string) [][]string {
	mp := map[[26]int][]string{}
	for _, s := range strs {
		k := [26]int{}
		for i := 0; i < len(s); i++ {
			k[s[i]-'a'] += 1
		}
		mp[k] = append(mp[k], s)
		fmt.Println(mp[k], s)
	}
	res := [][]string{}
	for _, v := range mp {
		res = append(res, v)
	}
	return res
}

func groupAnagrams2(str []string) [][]string {
	mp := map[[26]int][]string{}
	for _, s := range str {
		k := [26]int{}
		for i := 0; i < len(s); i++ {
			k[s[i]-'a'] += 1
		}
		mp[k] = append(mp[k], s)
	}
	result := [][]string{}
	for _, v := range mp {
		result = append(result, v)
	}
	return result
}
