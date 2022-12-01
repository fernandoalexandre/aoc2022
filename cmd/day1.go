/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"

	"github.com/spf13/cobra"
)

// day1Cmd represents the day1 command
var day1Cmd = &cobra.Command{
	Use:   "day1",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Day 1")
		log.Println("Parsing file...")
		inventory := readfile_p1("inputs/day1/input")
		elves := parse_elves_inventory(inventory)

		sort.Slice(elves[:], func(i, j int) bool {
			return elves[i].total_calories > elves[j].total_calories
		})

		log.Println("Part 1")
		part_1(elves)
		log.Println("Part 2")
		part_2(elves)
	},
}

type elf struct {
	calories       []int
	total_calories int
}

func readfile_p1(file_name string) []int {
	readFile, err := os.Open(file_name)

	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []int

	for fileScanner.Scan() {
		val, err := strconv.Atoi(fileScanner.Text())
		if err != nil {
			val = -1
		}
		fileLines = append(fileLines, val)
	}

	readFile.Close()

	return fileLines
}

func parse_elves_inventory(inventory []int) (elves []elf) {
	start_idx := 0
	curr_total := 0
	for idx, entry := range inventory {
		if entry == -1 {
			elves = append(elves, elf{inventory[start_idx:idx], curr_total})
			start_idx = idx + 1
			curr_total = 0
		} else {
			curr_total += entry
		}
	}
	return elves
}

func part_1(elves []elf) {
	log.Println(fmt.Sprintf("Most calories: %d", elves[0].total_calories))
}

func part_2(elves []elf) {
	total := elves[0].total_calories + elves[1].total_calories + elves[2].total_calories
	log.Println(fmt.Sprintf("Top 3 calories elves: %v", elves[0:3]))
	log.Println(fmt.Sprintf("Sum of top 3 calories: %d", total))
}

func init() {
	rootCmd.AddCommand(day1Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day1Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day1Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
