package main

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"strings"
	"time"
)

const (
	input = "ngcjuoqr"
)

func justHash(s string) string {
	h := md5.New()
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func getMd5(s string, i int, stretch bool, cache map[int]string) string {
	if hash, ok := cache[i]; ok {
		return hash
	}
	v := justHash(fmt.Sprintf("%v%d", s, i))
	if stretch {
		for i := 0; i < 2016; i++ {
			v = justHash(v)
		}
	}
	cache[i] = v
	return v
}

func substrToFind(s string) (string, error) {
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+1] && s[i+1] == s[i+2] {
			return strings.Repeat(string(s[i]), 5), nil
		}
	}
	return "", errors.New("no triple found")
}

func findKeyIndex(needed int, stretch bool, cache map[int]string) int {
	i := 0
	for found := 0; found < needed; {
		i++
		hash := getMd5(input, i, stretch, cache)
		if ss, err := substrToFind(hash); err == nil {
			for j := 1; j <= 1000; j++ {
				nextHash := getMd5(input, i+j, stretch, cache)
				if strings.Contains(nextHash, ss) {
					found++
					break
				}
			}
		}
	}
	return i
}

func part1() {
	fmt.Println("The index of the 64th key is:", findKeyIndex(64, false, map[int]string{}))
}

func part2() {
	fmt.Println("The index of the 64th key (with stretching) is:", findKeyIndex(64, true, map[int]string{}))
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}
