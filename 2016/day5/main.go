package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"time"
)

const (
	input = "reyedfim"
)

func getMd5(i int) string {
	h := md5.New()
	io.WriteString(h, fmt.Sprintf("%v%d", input, i))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func part1() {
	var answer string
	for i := 0; len(answer) < 8; i++ {
		v := getMd5(i)
		if v[:5] == "00000" {
			answer += string(v[5])
		}
	}
	fmt.Println("Answer:", answer)
}

func part2() {
	answer := []rune("********")
OUTER:
	for i := 0; ; i++ {
		v := getMd5(i)
		if v[:5] == "00000" {
			idx := v[5] - '0'
			if idx >= 0 && idx <= 7 {
				if ok := answer[idx]; ok == '*' {
					answer[idx] = rune(v[6])
					fmt.Println(string(answer))
				}
			}
			for _, r := range answer {
				if r == '*' {
					continue OUTER
				}
			}
			fmt.Println("Inner door password:", string(answer))
			break
		}
	}
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}
