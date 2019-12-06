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

func (o *object) NumOrbits() int {
	if o.orbits == nil {
		return 0
	}
	return o.orbits.NumOrbits() + 1
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
