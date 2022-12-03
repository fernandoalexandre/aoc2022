/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// day3Cmd represents the day3 command
var day3Cmd = &cobra.Command{
	Use:   "day3",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day3 called")
		fmt.Println("Testing part 1")
		/* test_bps := d3p1_parse_backpacks([]string{"vJrwpWtwJgWrhcsFMMfFFhFp",
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		"PmmdzqPrVvPwwTWBwg",
		"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
		"ttgJtRGJQctTZtZT",
		"CrZsJsPPZsGzwwsLwLmpwMDw"}) */
		// d3p1(test_bps)
		backpacks := d3p1_parse_backpacks(d3p1_readfile("inputs/day3/input"))
		d3p1(backpacks)
		// d3p2(test_bps)
		d3p2(backpacks)
	},
}

type Backpack struct {
	pocket_1 []int32 // first pocket with priority values
	pocket_2 []int32 // second pocket with priority values
	complete []int32
}

func d3p1_parse_backpacks(backpacks []string) (pbackpacks []Backpack) {
	log.Println("Parsing backpacks...")

	n_bps := len(backpacks)
	pbackpacks = make([]Backpack, n_bps)
	for idx, raw_bp := range backpacks {
		items := []int32(raw_bp)

		pbackpacks[idx] = Backpack{items[0 : len(items)/2], items[len(items)/2:], items}

		// log.Printf("Added backpack %v\n", pbackpacks[idx])
	}
	return pbackpacks
}

func d3p1_readfile(file_name string) []string {
	readFile, err := os.Open(file_name)

	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()

	return fileLines
}

func int32InSlice(a int32, list []int32) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func get_priority(item rune) int32 {
	if item >= int32('A') && item <= int32('Z') {
		return item%int32('A') + 27
	} else {
		return item%int32('a') + 1
	}
}

func d3p1(backpacks []Backpack) {
	log.Println("Part 1")

	total_common_priority := int32(0)

	for _, backpack := range backpacks {
		already_seen := make([]int32, len(backpack.pocket_1))
		for _, item := range backpack.pocket_1 {
			if int32InSlice(item, backpack.pocket_2) {
				if !int32InSlice(item, already_seen) {
					already_seen = append(already_seen, item)

					val := get_priority(item)

					total_common_priority += val
					// log.Printf("Backpack %d; Found common item %s|%d|%d\n", i, string(item), item, val)
					break
				}
			}
		}
	}

	log.Printf("Total common priority is: %d\n", total_common_priority)
}

func d3p2(backpacks []Backpack) {
	log.Println("Part 2")

	total_badge_priority := int32(0)

	// Foreach group
	for i := 0; i < len(backpacks)/3; i++ {
		curr_idx := i * 3
		for _, item := range backpacks[curr_idx].complete {
			if int32InSlice(item, backpacks[curr_idx+1].complete) && int32InSlice(item, backpacks[curr_idx+2].complete) {
				val := get_priority(item)

				total_badge_priority += val
				// log.Printf("Group %d; Found badge %s|%d|%d\n", i, string(item), item, val)
				break
			}
		}
	}
	log.Printf("Total badge priority is: %d\n", total_badge_priority)
}

func init() {
	rootCmd.AddCommand(day3Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day3Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day3Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
