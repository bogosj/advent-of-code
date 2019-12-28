package cryostasis

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	combinations "github.com/mxschmitt/golang-combinations"

	"github.com/bogosj/advent-of-code/2019/fileinput"

	"github.com/bogosj/advent-of-code/2019/computer"
)

// Cryostasis represents the state of the starship.
type Cryostasis struct {
	c *computer.Computer
}

// New creates a new environment.
func New(p string) *Cryostasis {
	c := Cryostasis{}
	c.c = computer.New(p)
	return &c
}

func allItems() ([]string, [][]string) {
	all := []string{"asterisk",
		"ornament",
		"cake",
		"space heater",
		"festive hat",
		"semiconductor",
		"food ration",
		"sand",
	}
	perms := combinations.All(all)
	return all, perms
}

func buildTestAll() (ret [][]string) {
	all, perms := allItems()
	for _, perm := range perms {
		r := []string{}
		for _, item := range all {
			r = append(r, "drop "+item+"\n")
		}
		for _, item := range perm {
			r = append(r, "take "+item+"\n")
		}
		r = append(r, "west\n")
		ret = append(ret, r)
	}
	return
}

// Run kicks off the environment.
func (c *Cryostasis) Run() {
	in := make(chan int, 50000)
	out := c.c.Compute(in)
	var testAll [][]string
	var commands []string
	for {
		time.Sleep(time.Millisecond)
		select {
		case o := <-out:
			fmt.Print(string(o))
		default:
			if c.c.AwaitingInput {
				var text string
				var err error
				if len(testAll) > 0 && len(commands) == 0 {
					commands = testAll[0]
					testAll = testAll[1:]
				}
				if len(commands) > 0 {
					text = commands[0]
					commands = commands[1:]
				} else {
					reader := bufio.NewReader(os.Stdin)
					text, err = reader.ReadString('\n')
					if err != nil {
						panic(err)
					}
				}
				if text == "DUMP\n" {
					ioutil.WriteFile("dump.txt", []byte(c.c.Dumps()), 0644)
				}
				if text == "LOAD\n" {
					lines := fileinput.ReadLines("dump.txt")
					c.c.Loads(lines[0])
				}
				if text == "TESTALL\n" {
					testAll = buildTestAll()
				}
				for _, r := range text {
					in <- int(r)
				}
				time.Sleep(time.Millisecond)
			}
		}
	}
}
