package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type object struct {
	name   string
	orbits *object
}

func (o *object) String() string {
	return o.name
}

func (o *object) NumOrbits() int {
	return len(o.PathToCOM()) - 1
}

func (o *object) PathToCOM() []string {
	if o.orbits == nil {
		return []string{o.name}
	}
	return append([]string{o.name}, o.orbits.PathToCOM()...)
}

func getObj(objs map[string]*object, name string) *object {
	obj, ok := objs[name]
	if !ok {
		o := object{name: name}
		obj = &o
		objs[name] = obj
	}
	return obj
}

func establishOrbit(objs map[string]*object, name1, name2 string) {
	obj1 := getObj(objs, name1)
	obj2 := getObj(objs, name2)
	obj2.orbits = obj1
}

func sharedOrbit(path1, path2 []string) (string, int) {
	for i, o1 := range path1 {
		for j, o2 := range path2 {
			if o1 == o2 {
				return o1, (i + j - 2)
			}
		}
	}
	return "", 0
}

func main() {
	objs := map[string]*object{}
	for _, pair := range input() {
		establishOrbit(objs, pair[0], pair[1])
	}

	totalOrbits := 0
	for _, v := range objs {
		totalOrbits += v.NumOrbits()
	}
	fmt.Println(totalOrbits)

	youPath := objs["YOU"].PathToCOM()
	sanPath := objs["SAN"].PathToCOM()
	so, dist := sharedOrbit(youPath, sanPath)
	fmt.Println(youPath)
	fmt.Println(sanPath)
	fmt.Printf("%v: %v\n", so, dist)
}

func input() [][]string {
	var ret [][]string
	lines := strings.Split(rawinput(), "\n")
	for _, line := range lines {
		vals := strings.Split(line, ")")
		if len(vals) == 2 {
			ret = append(ret, vals)
		}
	}
	return ret
}

func rawinput() string {
	data, _ := ioutil.ReadFile("input.txt")
	return string(data)
}
