package main

import (
	"strings"
)

// uniqueWord returns how many  unique words are in the map
func uniqueWord(ws map[string]int) int {
	us := 0
	for _, v := range ws {
		if v == 1 {
			us = us + 1
		}
	}

	return us
}

// repeatedWord returns the most repeated word in the map
func repeatedWord(ws map[string]int) string {
	w := ""
	c := 0
	for k, v := range ws {
		if v > c {
			w = k
			c = v
		}
	}
	return w
}

func solve(s string) (int, string) {
	mws := make(map[string]int, 0)

	s = strings.ToLower(s)
	ws := strings.Fields(s)

	for _, v := range ws {
		_, ok := mws[v]
		if ok {
			mws[v] = mws[v] + 1
		} else {
			mws[v] = 1
		}
	}
	u := uniqueWord(mws)
	w := repeatedWord(mws)
	return u, w
}

func main() {

}
