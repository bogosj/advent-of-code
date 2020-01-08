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

var (
	hashes = map[int]string{}
)

func getMd5(i int) string {
	if hash, ok := hashes[i]; ok {
		return hash
	}
	h := md5.New()
	io.WriteString(h, fmt.Sprintf("%v%d", input, i))
	hashes[i] = fmt.Sprintf("%x", h.Sum(nil))
	return hashes[i]
}

func substrToFind(s string) (string, error) {
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+1] && s[i+1] == s[i+2] {
			return strings.Repeat(string(s[i]), 5), nil
		}
	}
	return "", errors.New("no triple found")
}

func findKeyIndex(needed int) int {
	i := 0
	for found := 0; found < needed; {
		i++
		hash := getMd5(i)
		if ss, err := substrToFind(hash); err == nil {
			for j := 1; j <= 1000; j++ {
				nextHash := getMd5(i + j)
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
	fmt.Println("The index of the 64th key is:", findKeyIndex(64))
}

func part2() {
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}
