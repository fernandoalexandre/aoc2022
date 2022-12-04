/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// day4Cmd represents the day4 command
var day4Cmd = &cobra.Command{
	Use:   "day4",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day4 called")
		test_input := []pair{create_pair("2-4,6-8"),
			create_pair("2-3,4-5"),
			create_pair("5-7,7-9"),
			create_pair("2-8,3-7"),
			create_pair("6-6,4-6"),
			create_pair("2-6,4-8"),
		}
		d4p1(test_input)
		input := d4p1_readfile("inputs/day4/input")
		d4p1(input)
		d4p2(test_input)
		d4p2(input)
	},
}

type interval struct {
	start  int
	finish int
}

type pair struct {
	first  interval
	second interval
}

func (a interval) get_length() int {
	return a.finish - a.start
}

func (a pair) has_contained() bool {
	var bigger, smaller interval
	if a.first.get_length() >= a.second.get_length() {
		bigger = a.first
		smaller = a.second
	} else {
		bigger = a.second
		smaller = a.first
	}

	return bigger.start <= smaller.start && bigger.finish >= smaller.finish
}

func (a pair) has_intersection() bool {
	result := a.has_contained()
	result = result || a.second.start <= a.first.start && a.first.start <= a.second.finish
	result = result || a.second.start <= a.first.finish && a.first.finish <= a.second.finish
	result = result || a.first.start <= a.second.start && a.second.start <= a.first.finish
	result = result || a.first.start <= a.second.finish && a.second.finish <= a.first.finish
	return result
}

func create_pair(input string) (p pair) {
	split_pair := strings.Split(input, ",")

	split_pair1 := strings.Split(split_pair[0], "-")
	split_pair2 := strings.Split(split_pair[1], "-")

	v1_0, _ := strconv.Atoi(split_pair1[0])
	v1_1, _ := strconv.Atoi(split_pair1[1])
	p.first = interval{v1_0, v1_1}

	v1_0, _ = strconv.Atoi(split_pair2[0])
	v1_1, _ = strconv.Atoi(split_pair2[1])
	p.second = interval{v1_0, v1_1}

	return p
}

func d4p1_readfile(file_name string) []pair {
	readFile, err := os.Open(file_name)

	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []pair

	for fileScanner.Scan() {
		fileLines = append(fileLines, create_pair(fileScanner.Text()))
	}

	readFile.Close()

	return fileLines
}

func d4p1(interval_list []pair) {
	fmt.Println("Part 1")

	total_contained := 0
	for _, c_pair := range interval_list {
		if c_pair.has_contained() {
			total_contained += 1
		}
	}

	fmt.Printf("Total contained intervals: %d\n", total_contained)
}

func d4p2(interval_list []pair) {
	fmt.Println("Part 2")

	total_intersect := 0
	for _, c_pair := range interval_list {
		if c_pair.has_intersection() {
			total_intersect += 1
		}
	}

	fmt.Printf("Total intersected intervals: %d\n", total_intersect)
}

func init() {
	rootCmd.AddCommand(day4Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day4Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day4Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
