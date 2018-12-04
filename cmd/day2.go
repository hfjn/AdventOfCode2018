package cmd

import (
	"fmt"
	"os"
	"unicode/utf8"

	"github.com/spf13/cobra"
)

const file = "examples/day2.txt"

var day2 = &cobra.Command{
	Use:   "day2",
	Short: "Second day of advent",
	Run:   find_neighbor,
}

func count_characters(input string) map[rune]int {
	characters := make(map[rune]int)
	for _, character := range input {
		characters[character]++
	}
	return characters
}

func find_twos_and_threes(input string) (int, int) {
	characters := count_characters(input)
	twos := 0
	threes := 0
	for _, value := range characters {
		if value == 2 {
			twos = 1
		}

		if value == 3 {
			threes = 1
		}
	}
	return twos, threes
}

func computeDistance(a, b string) int {
	if len(a) == 0 {
		return utf8.RuneCountInString(b)
	}

	if len(b) == 0 {
		return utf8.RuneCountInString(a)
	}

	if a == b {
		return 0
	}

	// We need to convert to []rune if the strings are non-ascii.
	// This could be avoided by using utf8.RuneCountInString
	// and then doing some juggling with rune indices.
	// The primary challenge is keeping track of the previous rune.
	// With a range loop, its not that easy. And with a for-loop
	// we need to keep track of the inter-rune width using utf8.DecodeRuneInString
	s1 := []rune(a)
	s2 := []rune(b)

	// swap to save some memory O(min(a,b)) instead of O(a)
	if len(s1) > len(s2) {
		s1, s2 = s2, s1
	}
	lenS1 := len(s1)
	lenS2 := len(s2)

	// init the row
	x := make([]int, lenS1+1)
	for i := 0; i <= lenS1; i++ {
		x[i] = i
	}

	// fill in the rest
	for i := 1; i <= lenS2; i++ {
		prev := i
		var current int

		for j := 1; j <= lenS1; j++ {

			if s2[i-1] == s1[j-1] {
				current = x[j-1] // match
			} else {
				current = min(min(x[j-1]+1, prev+1), x[j]+1)
			}
			x[j-1] = prev
			prev = current
		}
		x[lenS1] = prev
	}
	return x[lenS1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func find_neighbor(ccmd *cobra.Command, args []string) {
	lines := read_file(file)
	for _, line := range lines {
		for _, line2 := range lines {
			if computeDistance(line, line2) == 1 {
				fmt.Println(line)
				fmt.Println(line2)
				os.Exit(0)
			}
		}
	}

}

func checksum(ccmd *cobra.Command, args []string) {
	twos := 0
	threes := 0
	for _, line := range read_file(file) {
		two, three := find_twos_and_threes(line)
		twos = twos + two
		threes = threes + three
	}
	fmt.Println(twos * threes)
}

func init() {
	rootCmd.AddCommand(day2)
}
