package reindeerolympics

import (
	"fmt"
	"strings"

	"github.com/bogosj/advent-of-code/intmath"

	"github.com/bogosj/advent-of-code/fileinput"
)

type reindeer struct {
	name       string
	speed      int
	dist       int
	runTime    int
	restNeeded int
	flyTime    int
	restTime   int
	points     int
}

// Race represents a race of reindeer.
type Race struct {
	racers []*reindeer
}

//Vixen can fly 19 km/s for 7 seconds, but then must rest for 124 seconds.
func (r *Race) load(p string) {
	lines := fileinput.ReadLines(p)
	for _, line := range lines {
		f := strings.Fields(line)
		rd := reindeer{
			name:       f[0],
			speed:      intmath.Atoi(f[3]),
			flyTime:    intmath.Atoi(f[6]),
			restNeeded: intmath.Atoi(f[13]),
		}
		r.racers = append(r.racers, &rd)
	}
}

// New creates a new race
func New(p string) *Race {
	r := Race{}
	r.load(p)
	return &r
}

// Run runs the race for a specified time.
func (r *Race) Run(t int) {
	for i := 0; i < t; i++ {
		for _, rd := range r.racers {
			if rd.runTime < rd.flyTime {
				rd.dist += rd.speed
				rd.runTime++
				continue
			}
			if rd.restTime < rd.restNeeded-1 {
				rd.restTime++
			} else {
				rd.restTime = 0
				rd.runTime = 0
			}
		}
		r.givePointsToLeaders()
	}
}

func (r *Race) givePointsToLeaders() {
	d := 0
	for _, rd := range r.racers {
		if d < rd.dist {
			d = rd.dist
		}
	}
	for _, rd := range r.racers {
		if rd.dist == d {
			rd.points++
		}
	}
}

// Winner2 returns a message about who is the winner based on points.
func (r *Race) Winner2() string {
	var winner reindeer
	for _, rd := range r.racers {
		fmt.Printf("%v: %d\n", rd.name, rd.points)
		if rd.points > winner.points {
			winner = *rd
		}
	}
	return fmt.Sprintf("Winner is %v with points of %d", winner.name, winner.points)
}

// Winner returns a message about who is the winner.
func (r *Race) Winner() string {
	var winner reindeer
	for _, rd := range r.racers {
		fmt.Printf("%v: %d\n", rd.name, rd.dist)
		if rd.dist > winner.dist {
			winner = *rd
		}
	}
	return fmt.Sprintf("Winner is %v with distance of %d", winner.name, winner.dist)
}
