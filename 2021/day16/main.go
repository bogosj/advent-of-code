package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
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

func versionSum(message string, inBinary bool) (int, string) {
	versionVal := 0
	bits := message
	if !inBinary {
		bits = messageToBin(bits)
	}
	for len(bits) > 0 {
		if len(bits) < 4 || binToInt(bits) == 0 {
			return versionVal, bits
		}
		version := binToInt(bits[:3])
		bits = bits[3:]
		typeId := binToInt(bits[:3])
		bits = bits[3:]

		// value packet
		if typeId == 4 {
			_, remainder := literalPacket(bits)
			bits = remainder
			versionVal += version
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
			subversionVal, _ := versionSum(subPackets, true)
			versionVal += subversionVal
		} else {
			subPacketCount := binToInt(bits[:11])
			bits = bits[11:]
			for i := 0; i < subPacketCount; i++ {
				subversionVal, leftoverBits := versionSum(bits, true)
				versionVal += subversionVal
				bits = leftoverBits
			}
		}
	}
	return versionVal, bits
}

func part1(in []string) {
	fmt.Println(versionSum(in[0], false))
}

func part2(in []string) {
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
