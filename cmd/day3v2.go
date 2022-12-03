/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// day3v2Cmd represents the day3v2 command
var day3v2Cmd = &cobra.Command{
	Use:   "day3v2",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day3v2 called")
		fmt.Println("Testing part 1")
		/* test_bps := d3p1_parse_backpacks([]string{"vJrwpWtwJgWrhcsFMMfFFhFp",
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		"PmmdzqPrVvPwwTWBwg",
		"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
		"ttgJtRGJQctTZtZT",
		"CrZsJsPPZsGzwwsLwLmpwMDw"}) */
		// d3p1(test_bps)
		backpacks := d3v2p1_parse_backpacks(d3p1_readfile("inputs/day3/input"))
		d3v2p1(backpacks)
		// d3p2(test_bps)
		d3v2p2(backpacks)
	},
}

type Backpackv2 struct {
	pocket_1 map[rune]int32 // first pocket with priority values
	pocket_2 map[rune]int32 // second pocket with priority values
	complete map[rune]int32
}

func d3v2p1_parse_backpacks(backpacks []string) (pbackpacks []Backpackv2) {
	log.Println("Parsing backpacks...")

	n_bps := len(backpacks)
	pbackpacks = make([]Backpackv2, n_bps)
	for idx, raw_bp := range backpacks {
		items := []int32(raw_bp)
		pbackpacks[idx].pocket_1 = make(map[rune]int32)
		pbackpacks[idx].pocket_2 = make(map[rune]int32)
		pbackpacks[idx].complete = make(map[rune]int32)

		for i, item := range items {
			if i < len(items)/2 {
				pbackpacks[idx].pocket_1[item] = 0
			} else {
				pbackpacks[idx].pocket_2[item] = 0
			}
			pbackpacks[idx].complete[item] = 0
		}

		// log.Printf("Added backpack %v\n", pbackpacks[idx])
	}
	return pbackpacks
}

func d3v2p1(backpacks []Backpackv2) {
	log.Println("Part 1")

	total_common_priority := int32(0)

	for _, backpack := range backpacks {
		already_seen := make(map[int32]int32)
		for key := range backpack.pocket_1 {
			_, ok := backpack.pocket_2[key]
			if ok {
				_, as_ok := already_seen[key]
				if !as_ok {
					already_seen[key] = 0

					val := get_priority(key)

					total_common_priority += val
					// log.Printf("Backpack %d; Found common item %s|%d|%d\n", i, string(item), item, val)
					break
				}
			}
		}
	}

	log.Printf("Total common priority is: %d\n", total_common_priority)
}

func d3v2p2(backpacks []Backpackv2) {
	log.Println("Part 2")

	total_badge_priority := int32(0)

	// Foreach group
	for i := 0; i < len(backpacks)/3; i++ {
		curr_idx := i * 3
		for key := range backpacks[curr_idx].complete {
			_, ok_1 := backpacks[curr_idx+1].complete[key]
			_, ok_2 := backpacks[curr_idx+2].complete[key]
			if ok_1 && ok_2 {
				val := get_priority(key)

				total_badge_priority += val
				// log.Printf("Group %d; Found badge %s|%d|%d\n", i, string(item), item, val)
				break
			}
		}
	}
	log.Printf("Total badge priority is: %d\n", total_badge_priority)
}

func init() {
	rootCmd.AddCommand(day3v2Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day3v2Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day3v2Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
