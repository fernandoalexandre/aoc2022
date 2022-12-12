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

// day12Cmd represents the day12 command
var day12Cmd = &cobra.Command{
	Use:   "day12",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day12 called")

		test_map := d12_readfile("inputs/day12/input_example")
		actual_map := d12_readfile("inputs/day12/input")

		d12p1(test_map)
	},
}

type MapNode struct {
	value rune
	x, y  int
	links []*MapNode
}

func d12_get_str(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

func (a MapNode) d12_can_move(target MapNode) bool {
	return Abs(int(a.value)-int(target.value)) <= 1
}

func d12_parse_map(input_map []string) (output_map map[string]MapNode, start_node MapNode, dest_node MapNode) {
	output_map = make(map[string]MapNode)
	for y, line := range input_map {
		for x, slot := range line {
			val, found := output_map[d12_get_str(x, y)]
			if !found {
				val = MapNode{slot, x, y, []*MapNode{}}

				if val.value == 'S' {
					val.value = 'a'
					start_node = val
				} else if val.value == 'E' {

					dest_node = val
				}

				output_map[d12_get_str(x, y)] = val
			}

			if val.value == 'S' {
				val.value = 'a'
				start_node = val
			} else if val.value == 'E' {
				val.value = 'z'
				dest_node = val
			}

			// x y+1
			if y < len(input_map) {
				r_val, found := output_map[d12_get_str(x, y+1)]
				if !found {
					val = MapNode{slot, x, y + 1, []*MapNode{}}
					output_map[d12_get_str(x, y+1)] = val
				}

				if val.d12_can_move(r_val) {
					val.links = append(val.links, &r_val)
					r_val.links = append(r_val.links, &val)
				}
			}

			// x+1 y
			if x < len(line) {
				t_val, found := output_map[d12_get_str(x+1, y)]
				if !found {
					val = MapNode{slot, x + 1, y, []*MapNode{}}
					output_map[d12_get_str(x+1, y)] = val
				}

				if val.d12_can_move(t_val) {
					val.links = append(val.links, &t_val)
					t_val.links = append(t_val.links, &val)
				}
			}
		}
	}
	return output_map, start_node, dest_node
}

func d12_readfile(file_name string) (content []string) {
	readFile, err := os.Open(file_name)

	if err != nil {
		log.Fatal(err)
	}

	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		content = append(content, fileScanner.Text())
	}
	return content
}

func find_shortest_route() {

}

func d12p1(raw_map []string) {
	log.Println("Part 1")
	parsed_map, start, dest := d12_parse_map(raw_map)

}

func init() {
	rootCmd.AddCommand(day12Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day12Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day12Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
