package util

import (
	"bufio"
	"os"
)

func ReadFile(file string) ([]string, func() error) {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, f.Close
}
