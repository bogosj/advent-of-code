package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"time"
)

const (
	input = "iwrupvqb"
)

func getMd5(i int) string {
	h := md5.New()
	io.WriteString(h, fmt.Sprintf("%v%d", input, i))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func part1() {
	var i int
	for i = 0; ; i++ {
		v := getMd5(i)
		if v[0] == '0' && v[1] == '0' && v[2] == '0' && v[3] == '0' && v[4] == '0' {
			break
		}
	}
	fmt.Println("Answer:", i)
}

func part2() {
	var i int
	for i = 0; ; i++ {
		v := getMd5(i)
		if v[0] == '0' && v[1] == '0' && v[2] == '0' && v[3] == '0' && v[4] == '0' && v[5] == '0' {
			break
		}
	}
	fmt.Println("Answer:", i)
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}
