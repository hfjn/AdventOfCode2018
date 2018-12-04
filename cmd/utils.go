package cmd

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

func read_file(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines
}

func mapSubexpNames(m, n []string) map[string]int {
	m, n = m[1:], n[1:]
	r := make(map[string]int, len(m))
	for i, _ := range n {
		r[n[i]], _ = strconv.Atoi(m[i])
	}
	return r
}

func parse_string(regex, line string) map[string]int {
	r, _ := regexp.Compile(regex)
	m := r.FindStringSubmatch(line)
	n := r.SubexpNames()
	return mapSubexpNames(m, n)
}
