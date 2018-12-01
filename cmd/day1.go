package cmd

import (
	"bufio"
	"fmt"
	"log"
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

	numbers := read_file(filename)

	current := 0
	occurrences := make(map[int]int)

	found := false

	for !found {
		for _, number := range numbers {
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

func read_file(filename string) []int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var numbers []int

	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
			os.Exit(2)
		}
		numbers = append(numbers, number)
	}

	return numbers
}

func init() {
	rootCmd.AddCommand(day1)
}
