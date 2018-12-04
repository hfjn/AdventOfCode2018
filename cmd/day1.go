package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

const check_occurences = true
const filename = "examples/day1.txt"

var day1 = &cobra.Command{
	Use:   "day1",
	Short: "First day of advent",
	Run:   calculate,
}

func calculate(ccmd *cobra.Command, args []string) {

	lines := read_file(filename)

	current := 0
	occurrences := make(map[int]int)

	found := false

	for !found {
		for _, line := range lines {
			number, _ := strconv.Atoi(line)
			current = current + number
			if check_occurences {
				occurrences[current]++
				if occurrences[current] > 1 {
					fmt.Printf("%s%d", "first duplicate occurence is: ", current)
					os.Exit(0)
				}
			}
		}
	}

	fmt.Println(current)

}

func init() {
	rootCmd.AddCommand(day1)
}
