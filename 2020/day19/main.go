package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

type rule struct {
	children [][]int
	char     string
}

type message struct {
	rules []int
	msg   string
}

func (m message) expand(rules map[int]rule, prefixes map[string]bool) (ret []message) {
	for len(m.rules) > 0 {
		r := rules[m.rules[0]]
		if r.char == "" {
			break
		}
		m.msg += r.char
		m.rules = m.rules[1:]
	}

	if !prefixes[m.msg] {
		return nil
	}

	if len(m.rules) == 0 {
		return []message{m}
	}

	r := rules[m.rules[0]]
	for _, cg := range r.children {
		nm := message{}
		nm.rules = append(nm.rules, cg...)
		nm.rules = append(nm.rules, m.rules[1:]...)
		nm.msg = m.msg
		ret = append(ret, nm)
	}

	return
}

func generateMessages(rules map[int]rule, prefixes map[string]bool) (ret []message) {
	rootMsg := message{rules: rules[0].children[0]}
	msgs := []message{rootMsg}
	for len(msgs) > 0 {
		msg := msgs[0]
		msgs = msgs[1:]

		if len(msg.rules) == 0 {
			ret = append(ret, msg)
			continue
		}

		msgs = append(msgs, msg.expand(rules, prefixes)...)
	}
	return ret
}

func validPrefixes(msgs []string) map[string]bool {
	ret := map[string]bool{"": true}
	for _, msg := range msgs {
		for len(msg) > 0 {
			ret[msg] = true
			msg = msg[:len(msg)-1]
		}
	}
	return ret
}

func part1(rules map[int]rule, msgs []string) {
	count := 0
	prefixes := validPrefixes(msgs)
	allPossible := generateMessages(rules, prefixes)
	for _, msg := range msgs {
		for _, p := range allPossible {
			if msg == p.msg {
				count++
			}
		}
	}
	fmt.Printf("There are %v valid messages\n", count)
}

func part2(rules map[int]rule, msgs []string) {
}

func main() {
	rules, msgs := input()
	start := time.Now()
	part1(rules, msgs)
	fmt.Println("Part 1 done in", time.Since(start))
	rules, msgs = input()
	start = time.Now()
	part2(rules, msgs)
	fmt.Println("Part 2 done in", time.Since(start))
}

func input() (map[int]rule, []string) {
	rules := map[int]rule{}

	lines := fileinput.ReadLines("input.txt")
	for idx, line := range lines {
		if len(line) == 0 {
			return rules, lines[idx+1:]
		}

		f := strings.Fields(line)
		id := intmath.Atoi(strings.Replace(f[0], ":", "", 1))
		r := rule{}

		children := []int{}
		for _, token := range f[1:] {
			switch token {
			case "|":
				r.children = append(r.children, children)
				children = []int{}
			case `"a"`, `"b"`:
				r.char = strings.Replace(token, `"`, "", 2)
			default:
				children = append(children, intmath.Atoi(token))
			}
		}
		r.children = append(r.children, children)
		rules[id] = r
	}
	return nil, nil
}
