package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const file_name = "examples/day3.txt"

var day3 = &cobra.Command{
	Use:   "day3",
	Short: "Second day of advent",
	Run:   solve,
}

func solve(ccmd *cobra.Command, args []string) {
	lines := read_file(file_name)
	matrix := make(map[string][]int)
	var final_id int
	for _, line := range lines {
		regex := `^#(?P<id>\d+)\s@\s(?P<x>\d+),(?P<y>\d+):\s(?P<width>\d+)x(?P<height>\d+)$`
		parsed_values := parse_string(regex, line)
		right := parsed_values["x"] + parsed_values["width"]
		bottom := parsed_values["y"] + parsed_values["height"]
		id := parsed_values["id"]
		for x := parsed_values["x"]; x < right; x++ {
			for y := parsed_values["y"]; y < bottom; y++ {
				coord := fmt.Sprintf("%d%s%d", x, "_", y)

				matrix[coord] = append(matrix[coord], id)
			}
		}

	}

	fmt.Println(final_id)

	var dupes int
	duplicate_ids := make(map[int]bool)
	singular_ids := make(map[int]bool)
	for _, v := range matrix {

		if len(v) > 1 {
			for _, id := range v {
				duplicate_ids[id] = true
			}
			dupes++
		}
		if len(v) == 1 {
			for _, id := range v {
				singular_ids[id] = true
			}
		}
	}

	for k, v := range singular_ids {
		if v && !duplicate_ids[k] {
			fmt.Println(k)
		}
	}

	fmt.Println(dupes)
}

func init() {
	rootCmd.AddCommand(day3)
}
