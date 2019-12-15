package fileinput

import (
	"io/ioutil"
	"strings"
)

// ReadLines accepts a file path and returns a slice of strings of lines in the file.
func ReadLines(n string) []string {
	data, _ := ioutil.ReadFile(n)
	full := strings.TrimSpace(string(data))
	return strings.Split(full, "\n")
}
