package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

const (
	start = ".#./..#/###"
)

func splitImage(img [][]rune, factor int) (ret [][][]rune) {
	for len(img) > 0 {
		for len(img[0]) > 0 {
			var subImg [][]rune
			for i := 0; i < factor; i++ {
				subImg = append(subImg, img[i][0:factor])
				img[i] = img[i][factor:]
			}
			ret = append(ret, subImg)
		}
		img = img[factor:]
	}
	return
}

func mergeImages(subImgs [][][]rune) (ret [][]rune) {
	factor := intmath.Sqrt(len(subImgs))
	for len(subImgs) > 0 {
		for len(subImgs[0]) > 0 {
			var newRow []rune
			for i := 0; i < factor; i++ {
				newRow = append(newRow, subImgs[i][0]...)
				subImgs[i] = subImgs[i][1:]
			}
			ret = append(ret, newRow)
		}
		subImgs = subImgs[factor:]
	}
	return
}

func step(img [][]rune, transforms map[string]string) [][]rune {
	var factor int
	if len(img) == 2 || len(img) == 3 {
		return to2D(transforms[to1D(img)])
	} else if len(img)%2 == 0 {
		factor = 2
	} else if len(img)%3 == 0 {
		factor = 3
	}
	var newSubImgs [][][]rune
	subImgs := splitImage(img, factor)
	for _, subImg := range subImgs {
		newSubImgs = append(newSubImgs, step(subImg, transforms))
	}
	return mergeImages(newSubImgs)
}

func printImg(img [][]rune) {
	for _, row := range img {
		for _, cell := range row {
			fmt.Print(string(cell))
		}
		fmt.Println()
	}
}

func part1() {
	t := transforms()
	img := to2D(start)
	for i := 0; i < 5; i++ {
		img = step(img, t)
	}
	var c int
	for _, r := range to1D(img) {
		if r == '#' {
			c++
		}
	}
	fmt.Printf("There are %d pixels on\n", c)
}

func part2() {
	t := transforms()
	img := to2D(start)
	for i := 0; i < 18; i++ {
		img = step(img, t)
	}
	var c int
	for _, r := range to1D(img) {
		if r == '#' {
			c++
		}
	}
	fmt.Printf("There are %d pixels on\n", c)
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}

func input() (ret [][]string) {
	for _, line := range fileinput.ReadLines("input.txt") {
		f := strings.Fields(line)
		ret = append(ret, []string{f[0], f[2]})
	}
	return
}

func to2D(from string) (ret [][]rune) {
	for _, row := range strings.Split(from, "/") {
		ret = append(ret, []rune(row))
	}
	return
}

func to1D(from [][]rune) string {
	var retA []string
	for _, row := range from {
		retA = append(retA, string(row))
	}
	return strings.Join(retA, "/")
}

func rotate(from [][]rune) (ret [][]rune) {
	if len(from) == 2 {
		ret = append(ret, []rune{from[1][0], from[0][0]})
		ret = append(ret, []rune{from[1][1], from[0][1]})
	} else {
		ret = append(ret, []rune{from[2][0], from[1][0], from[0][0]})
		ret = append(ret, []rune{from[2][1], from[1][1], from[0][1]})
		ret = append(ret, []rune{from[2][2], from[1][2], from[0][2]})
	}
	return
}

func flipX(from [][]rune) (ret [][]rune) {
	if len(from) == 2 {
		ret = append(ret, []rune{from[0][1], from[0][0]})
		ret = append(ret, []rune{from[1][1], from[1][0]})
	} else {
		ret = append(ret, []rune{from[0][2], from[0][1], from[0][0]})
		ret = append(ret, []rune{from[1][2], from[1][1], from[1][0]})
		ret = append(ret, []rune{from[2][2], from[2][1], from[2][0]})
	}
	return
}

func flipY(from [][]rune) (ret [][]rune) {
	if len(from) == 2 {
		ret = append(ret, []rune{from[1][0], from[1][1]})
		ret = append(ret, []rune{from[0][0], from[0][1]})
	} else {
		ret = append(ret, []rune{from[2][0], from[2][1], from[2][2]})
		ret = append(ret, []rune{from[1][0], from[1][1], from[1][2]})
		ret = append(ret, []rune{from[0][0], from[0][1], from[0][2]})
	}
	return
}

func makeTransforms(from, to string, t map[string]string) {
	curr := to2D(from)
	for i := 0; i < 4; i++ {
		t[to1D(curr)] = to
		curr = flipX(curr)
		t[to1D(curr)] = to
		curr = flipX(curr)
		curr = flipY(curr)
		t[to1D(curr)] = to
		curr = flipY(curr)
		curr = rotate(curr)
	}
}

func transforms() map[string]string {
	ret := map[string]string{}
	for _, r := range input() {
		makeTransforms(r[0], r[1], ret)
	}
	return ret
}
