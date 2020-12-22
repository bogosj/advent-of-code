package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/albrow/stringset"
	"github.com/bogosj/advent-of-code/fileinput"
)

var (
	foodReg = regexp.MustCompile(`(.*) \(contains (.*),?\)`)
)

type food struct {
	data string
}

func (f food) ingredients() []string {
	return strings.Fields(foodReg.FindStringSubmatch(f.data)[1])
}

func (f food) allergens() []string {
	return strings.Split(foodReg.FindStringSubmatch(f.data)[2], ", ")
}

func getAllergenMap(in []food) map[string]string {
	allAllergens := stringset.New()
	for _, f := range in {
		allAllergens.Add(f.allergens()...)
	}

	allergenMap := map[string]string{}

	for len(allAllergens) > 0 {
		for a := range allAllergens {
			ingredients := []stringset.Set{}
			for _, f := range in {
				currrentAllergens := stringset.NewFromSlice(f.allergens())
				if currrentAllergens.Contains(a) {
					currentIngredients := stringset.NewFromSlice(f.ingredients())
					for k := range allergenMap {
						currentIngredients.Remove(k)
					}
					ingredients = append(ingredients, currentIngredients)
				}
			}

			s1 := ingredients[0]
			for _, s2 := range ingredients[1:] {
				s1 = stringset.Intersect(s1, s2)
			}
			if len(s1) == 1 {
				allergenMap[s1.Slice()[0]] = a
				allAllergens.Remove(a)
				break
			}
		}
	}
	return allergenMap
}

func part1(in []food) {
	allergenMap := getAllergenMap(in)
	count := 0
	for _, f := range in {
		for _, i := range f.ingredients() {
			if _, ok := allergenMap[i]; !ok {
				count++
			}
		}
	}
	fmt.Printf("The non-allergen ingredients appear %v times\n", count)
}

func part2(in []food) {
	allergenMap := getAllergenMap(in)
	inv := map[string]string{}
	allergens := []string{}
	for k, v := range allergenMap {
		inv[v] = k
		allergens = append(allergens, v)
	}
	sort.Strings(allergens)
	for _, v := range allergens {
		fmt.Printf("%v,", inv[v])
	}
	fmt.Println()
}

func main() {
	in := input()
	start := time.Now()
	part1(in)
	fmt.Println("Part 1 done in", time.Since(start))
	in = input()
	start = time.Now()
	part2(in)
	fmt.Println("Part 2 done in", time.Since(start))
}

func input() []food {
	ret := []food{}

	for _, line := range fileinput.ReadLines("input.txt") {
		ret = append(ret, food{data: line})
	}

	return ret
}
