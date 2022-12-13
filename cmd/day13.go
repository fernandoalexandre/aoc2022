/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/spf13/cobra"
)

// day13Cmd represents the day13 command
var day13Cmd = &cobra.Command{
	Use:   "day13",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day13 called")

		test_content := d13_readfile("inputs/day13/input_example")
		content := d13_readfile("inputs/day13/input")

		d13p1(test_content)
		d13p1(content)

		d13p2(test_content)
		d13p2(content)

	},
}

func d13_readfile(file_name string) (content [][]string) {
	readFile, err := os.Open(file_name)

	if err != nil {
		log.Fatal(err)
	}

	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	curr_pair := []string{}
	for fileScanner.Scan() {
		if fileScanner.Text() != "" {
			curr_pair = append(curr_pair, fileScanner.Text())
		} else {
			content = append(content, curr_pair)
			curr_pair = []string{}
		}
	}
	content = append(content, curr_pair)
	return content
}

func d13_is_ordered(lleft, lright []any) (is_ordered int) {
	idx := 0
	for i := 0; i < len(lleft); i++ {
		if i == len(lright) {
			return 0
		}

		is_ordered := -1
		if left, ok := lleft[idx].(float64); ok { // first is float64
			if right, ok := lright[idx].(float64); ok { // first is float64 && second is float64
				if left > right {
					return 0
				} else if left < right {
					return 1
				}
			} else if right, ok := lright[idx].([]any); ok { // first is float64 && second is []any
				is_ordered = d13_is_ordered([]any{left}, right)
			}
		} else if left, ok := lleft[idx].([]any); ok { // first is []any
			if right, ok := lright[idx].([]any); ok { // first is []any && second is []any
				is_ordered = d13_is_ordered(left, right)
			} else if right, ok := lright[idx].(float64); ok { // first is []any && second is float64
				is_ordered = d13_is_ordered(left, []any{right})
			}
		}
		if is_ordered != -1 {
			return is_ordered
		}
		idx++
	}

	if len(lleft) < len(lright) {
		return 1
	}

	return -1
}

func d13_less(i int, j int) bool {
	return d13_is_ordered(flat_list[i], flat_list[j]) == 1
}

func d13p1(content [][]string) {
	log.Println("Part 1")

	result := 0
	for idx, pair := range content {

		var left, right []any
		json.Unmarshal([]byte(pair[0]), &left)
		json.Unmarshal([]byte(pair[1]), &right)

		i_o := d13_is_ordered(left, right)
		if i_o == 1 || i_o == -1 {
			result += idx + 1
		}
	}
	log.Printf("Result: %d", result)
}

var flat_list [][]any

func d13p2(content [][]string) {
	log.Println("Part 2")

	flat_list = [][]any{}
	// Create a flat slice
	for _, pair := range content {
		var left, right []any
		json.Unmarshal([]byte(pair[0]), &left)
		json.Unmarshal([]byte(pair[1]), &right)

		flat_list = append(flat_list, left)
		flat_list = append(flat_list, right)
	}

	// Add the two new dividers
	var f, s []any
	json.Unmarshal([]byte("[[2]]"), &f)
	json.Unmarshal([]byte("[[6]]"), &s)
	flat_list = append(flat_list, f)
	flat_list = append(flat_list, s)

	// Sort it
	sort.Slice(flat_list, d13_less)

	result := 1
	for idx, entry := range flat_list {
		val, _ := json.Marshal(entry)
		str := string(val[:])
		if str == "[[2]]" || str == "[[6]]" {
			result *= idx + 1
		}
	}

	log.Printf("Result: %d", result)
}

func init() {
	rootCmd.AddCommand(day13Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day13Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day13Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
