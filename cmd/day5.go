package cmd

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

var day5 = &cobra.Command{
	Use:   "day5",
	Short: "Second day of advent",
	Run:   solve_day5,
}

func removeCharacters(input string, characters string) string {
	filter := func(r rune) rune {
		if strings.IndexRune(characters, r) < 0 {
			return r
		}
		return -1
	}
	return strings.Map(filter, input)

}

func solve_day5(ccmd *cobra.Command, args []string) {
	file_name := "examples/day5.txt"
	lines := read_file(file_name)
	input := strings.TrimSpace(lines[0])

	reg, err := regexp.Compile("[^a-zA-Z]+")
	if err != nil {
		log.Fatal(err)
	}
	input = reg.ReplaceAllString(input, "")

	// Part 1
	fmt.Println("Part 1:")
	fmt.Println(len(resolve("", input)))

	// Part 2
	result := make(map[string]int)
	for _, character := range "abcdefghijklmnopqrstuvwxyz" {
		characters := string(character) + strings.ToUpper(string(character))
		new_input := removeCharacters(input, characters)
		result[string(character)] = len(resolve("", new_input))
	}
	fmt.Println("Part 2:")
	fmt.Println(result)
}

func resolve(before, after string) string {
	// fmt.Println(before + " - " + after)
	if len(after) == 0 {
		return before
	}

	if len(before) == 0 {
		return resolve(string(after[0]), after[1:])
	}

	last_id := len(before) - 1
	left_character := string(before[last_id])
	right_character := string(after[0])
	if left_character != right_character && (strings.ToUpper(left_character) == right_character || left_character == strings.ToUpper(right_character)) {
		return resolve(before[:last_id], after[1:])
	}
	return resolve(before+string(after[0]), after[1:])

}

func init() {
	rootCmd.AddCommand(day5)
}
