package cmd

import (
	"fmt"
	"os"
	"regexp"
	"time"

	"github.com/spf13/cobra"
)

var day4 = &cobra.Command{
	Use:   "day4",
	Short: "Second day of advent",
	Run:   solve_day4,
}

func solve_day4(ccmd *cobra.Command, args []string) {
	file_name := "examples/day4.txt"
	lines := read_file(file_name)

	time_regex, _ := regexp.Compile(`\[([\d]+-[\d]+-[\d]+\s[\d]+:[\d]+)\]`)
	guard_regex, _ := regexp.Compile(`Guard\s#([\d]+)`)
	wake_regex, _ := regexp.Compile(`wakes\sup`)
	sleep_regex, _ := regexp.Compile(`falls\sasleep`)

	time_layout := "2006-01-02 15:04"
	guard := ""

	guards := make(map[string]int)
	best_minutes := make(map[string][]int)
	var start_timestamp time.Time

	for _, line := range lines {
		m := time_regex.FindStringSubmatch(line)
		timeStamp, err := time.Parse(time_layout, m[1])

		if err != nil {
			fmt.Println(err)
			os.Exit(1)

		}
		m = guard_regex.FindStringSubmatch(line)
		if len(m) > 1 {
			guard = m[1]
			if _, ok := best_minutes[guard]; !ok {
				best_minutes[guard] = make([]int, 60)
			}
		}

		m = sleep_regex.FindStringSubmatch(line)
		if len(m) > 0 {
			start_timestamp = timeStamp
		}

		m = wake_regex.FindStringSubmatch(line)
		if len(m) > 0 {
			guards[guard] = guards[guard] + int(timeStamp.Sub(start_timestamp).Minutes())
			if start_timestamp.Day() < timeStamp.Day() {
				for minute := 0; minute <= int(timeStamp.Minute()); minute++ {
					best_minutes[guard][minute]++
				}
			} else {
				for minute := int(start_timestamp.Minute()); minute <= int(timeStamp.Minute()); minute++ {
					best_minutes[guard][minute]++
				}
			}
		}
	}

	max_minutes := 0
	var best_guard string

	// Part 1
	for guard, minutes := range guards {
		// fmt.Printf("%s%s%d", guard, " - ", minutes)
		if minutes >= max_minutes {
			max_minutes = minutes
			best_guard = guard
		}
	}

	max_minutes = 0
	longest_minute := 0

	for k, v := range best_minutes[best_guard] {
		// fmt.Printf("%d%s%d", k, " - ", v)
		if v >= max_minutes {
			max_minutes = v
			longest_minute = k
		}
	}
	fmt.Println("Part 1")
	fmt.Println(best_guard)
	fmt.Println(longest_minute)

	// Part 2
	max_amount_single_minute := 0
	perfect_minute := 0

	for guard, minutes := range best_minutes {
		for minute := 0; minute < len(minutes); minute++ {
			amount := minutes[minute]
			if amount > max_amount_single_minute {
				perfect_minute = minute
				max_amount_single_minute = amount
				best_guard = guard
			}
		}
	}

	fmt.Println("Part 2")
	fmt.Println(best_guard)
	fmt.Println(max_amount_single_minute)
	fmt.Println(perfect_minute)
}

func init() {
	rootCmd.AddCommand(day4)
}
