package knothash

import (
	"fmt"
)

func makeList() []int {
	out := make([]int, 256)
	for i := 0; i < 256; i++ {
		out[i] = i
	}
	return out
}

func reverse(in []int, idx, length int) {
	l := len(in)
	for left, right := idx, idx+length-1; left < right; left, right = left+1, right-1 {
		in[left%l], in[right%l] = in[right%l], in[left%l]
	}
}

func rotateList(in []int, s string) []int {
	var idx, skipSize int
	for i := 0; i < 64; i++ {
		for _, length := range s {
			l := int(length)
			reverse(in, idx, l)
			idx = (idx + skipSize + l) % len(in)
			skipSize++
		}
	}
	return in
}

func Hash(s string) (ret string) {
	in := rotateList(makeList(), s)
	for len(in) > 0 {
		part := in[0:16]
		in = in[16:]
		v := part[0]
		for i := 1; i < len(part); i++ {
			v ^= part[i]
		}
		n := fmt.Sprintf("%x", v)
		if len(n) == 1 {
			n = "0" + n
		}
		ret += n
	}
	return
}
