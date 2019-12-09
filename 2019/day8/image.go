package main

type layer struct {
	l [][]int
}

func (l *layer) store(in []int) {
	for len(in) != 0 {
		l.l = append(l.l, in[:25])
		in = in[25:]
	}
}

func (l *layer) numZeros() int {
	digits := l.digits()
	return digits[0]
}

func (l *layer) merge(o layer) layer {
	merged := layer{}
	for i, row := range l.l {
		merged.l = append(merged.l, row)
		for j := range row {
			if merged.l[i][j] == 2 {
				merged.l[i][j] = o.l[i][j]
			}
		}
	}
	return merged
}

func (l *layer) String() (s string) {
	for _, row := range l.l {
		for _, pixel := range row {
			if pixel == 1 {
				s += "*"
			} else {
				s += " "
			}
		}
		s += "\n"
	}
	return
}

func (l *layer) digits() map[int]int {
	digits := map[int]int{}
	for _, row := range l.l {
		for _, pixel := range row {
			digits[pixel]++
		}
	}
	return digits
}

type image struct {
	layers []layer
}

func (i *image) parse(in []int) {
	for len(in) != 0 {
		l := layer{}
		l.store(in[:25*6])
		i.layers = append(i.layers, l)
		in = in[25*6:]
	}
}

func (i *image) flatten() *layer {
	merged := i.layers[0].merge(i.layers[1])
	for r := 2; r < len(i.layers); r++ {
		merged = merged.merge(i.layers[r])
	}
	return &merged
}
