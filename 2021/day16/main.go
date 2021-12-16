package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

func binToInt(bin string) int {
	i, _ := strconv.ParseInt(bin, 2, 64)
	return int(i)
}

func hexToBin(hex string) string {
	ui, _ := strconv.ParseUint(hex, 16, 64)
	format := fmt.Sprintf("%%0%db", len(hex)*4)
	return fmt.Sprintf(format, ui)
}

func messageToBin(message string) (ret string) {
	for _, char := range message {
		ret += hexToBin(string(char))
	}
	return
}

func literalPacket(bits string) (int, string) {
	valueBits := ""
	for {
		chunk := bits[:5]
		bits = bits[5:]
		valueBits += chunk[1:]

		if chunk[:1] == "0" {
			break
		}
	}
	return binToInt(valueBits), bits
}

func processSubValues(opType int, values []int) int {
	switch opType {
	case 0:
		return intmath.Sum(values...)
	case 1:
		return intmath.Product(values...)
	case 2:
		return intmath.Min(values...)
	case 3:
		return intmath.Max(values...)
	case 5:
		if values[0] > values[1] {
			return 1
		}
		return 0
	case 6:
		if values[0] < values[1] {
			return 1
		}
		return 0
	case 7:
		if values[0] == values[1] {
			return 1
		}
		return 0
	}
	return 0
}

func versionAndValue(message string, inBinary bool, packetCount int) (int, []int, string) {
	versionVal := 0
	values := []int{}
	bits := message
	if !inBinary {
		bits = messageToBin(bits)
	}
	for len(bits) > 0 {
		if packetCount > 0 && len(values) == packetCount {
			return versionVal, values, bits
		}
		if len(bits) < 4 || binToInt(bits) == 0 {
			return versionVal, values, bits
		}
		version := binToInt(bits[:3])
		bits = bits[3:]
		typeId := binToInt(bits[:3])
		bits = bits[3:]

		// value packet
		if typeId == 4 {
			value, remainder := literalPacket(bits)
			bits = remainder
			versionVal += version
			values = append(values, value)
			continue
		}

		versionVal += version

		// operator packet
		lengthType := binToInt(bits[:1])
		bits = bits[1:]
		if lengthType == 0 {
			length := binToInt(bits[:15])
			bits = bits[15:]
			subPackets := bits[:length]
			bits = bits[length:]
			subversionVal, subValues, _ := versionAndValue(subPackets, true, 0)
			values = append(values, processSubValues(typeId, subValues))
			versionVal += subversionVal
		} else {
			packetCount := binToInt(bits[:11])
			bits = bits[11:]
			subversionVal, subValues, leftoverBits := versionAndValue(bits, true, packetCount)
			values = append(values, processSubValues(typeId, subValues[:packetCount]))
			versionVal += subversionVal
			bits = leftoverBits
		}
	}
	return versionVal, values, bits
}

func part1(in []string) {
	versionSum, _, _ := versionAndValue(in[0], false, 0)
	fmt.Println("Part 1 answer:", versionSum)
}

func part2(in []string) {
	_, value, _ := versionAndValue(in[0], false, 0)
	fmt.Println("Part 2 answer:", value[0])
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

func input() []string {
	return fileinput.ReadLines("input.txt")
}
