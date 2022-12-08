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

// day8Cmd represents the day8 command
var day8Cmd = &cobra.Command{
	Use:   "day8",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day8 called")
		test_tree_map := []string{"30373",
			"25512",
			"65332",
			"33549",
			"35390"}

		tree_map := d8_convert_map(d8_readfile("inputs/day8/input"))
		d8p1(d8_convert_map(test_tree_map))
		d8p1(tree_map)

		d8p2(d8_convert_map(test_tree_map))
		d8p2(tree_map)
	},
}

func d8_readfile(file_name string) (content []string) {
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

func d8_convert_map(raw_map []string) (tree_map [][]int) {
	log.Println("Converting map...")
	for _, line := range raw_map {
		tree_map_line := []int{}
		for _, char := range line {
			tree_map_line = append(tree_map_line, int(char))
		}

		tree_map = append(tree_map, tree_map_line)
	}
	return tree_map
}

func d8p1_is_visible(tree_map [][]int, i int, j int) (is_visible bool) {
	is_visible = false
	// top col
	is_visible_top := true
	//log.Printf("Processing %d, %d (%d, %d)", i, j, len(tree_map), len(tree_map[0]))
	for row := i - 1; row >= 0; row-- {
		if tree_map[i][j] <= tree_map[row][j] {
			is_visible_top = false
			break
		}
	}

	// bottom col
	is_visible_bottom := true
	for row := i + 1; row < len(tree_map); row++ {
		if tree_map[i][j] <= tree_map[row][j] {
			is_visible_bottom = false
			break
		}
	}

	// left row
	is_visible_left := true
	for col := j - 1; col >= 0; col-- {
		if tree_map[i][j] <= tree_map[i][col] {
			is_visible_left = false
			break
		}
	}

	// right row
	is_visible_right := true
	for col := j + 1; col < len(tree_map[0]); col++ {
		if tree_map[i][j] <= tree_map[i][col] {
			is_visible_right = false
			break
		}
	}
	return is_visible_top || is_visible_bottom || is_visible_left || is_visible_right
}

func d8p1(tree_map [][]int) {
	log.Println("Part 1")

	num_visible_trees := 0

	for i := 0; i < len(tree_map); i++ {
		for j := 0; j < len(tree_map[0]); j++ {
			// Is edge?
			if i == 0 || i == len(tree_map)-1 || j == 0 || j == len(tree_map[0])-1 || d8p1_is_visible(tree_map, i, j) {
				num_visible_trees += 1
			}
		}
	}
	log.Printf("Number of visible trees: %d", num_visible_trees)
}

func d8p2_scenic_score(tree_map [][]int, i int, j int) int {
	// top col
	scenic_score_top := 0
	//log.Printf("Processing %d, %d (%d, %d)", i, j, len(tree_map), len(tree_map[0]))
	for row := i - 1; row >= 0; row-- {
		scenic_score_top = i - row
		if tree_map[i][j] <= tree_map[row][j] {
			break
		}
	}

	// bottom col
	scenic_score_bottom := 0
	for row := i + 1; row < len(tree_map); row++ {
		scenic_score_bottom = row - i
		if tree_map[i][j] <= tree_map[row][j] {
			break
		}
	}

	// left row
	scenic_score_left := 0
	for col := j - 1; col >= 0; col-- {
		scenic_score_left = j - col
		if tree_map[i][j] <= tree_map[i][col] {
			break
		}
	}

	// right row
	scenic_score_right := 0
	for col := j + 1; col < len(tree_map[0]); col++ {
		scenic_score_right = col - j
		if tree_map[i][j] <= tree_map[i][col] {
			break
		}
	}
	log.Printf("%d * %d * %d * %d (%d, %d)", scenic_score_top, scenic_score_bottom, scenic_score_left, scenic_score_right, i, j)
	return scenic_score_top * scenic_score_bottom * scenic_score_left * scenic_score_right
}

func d8p2(tree_map [][]int) {
	log.Println("Part 2")

	scenic_score := 0

	for i := 1; i < len(tree_map)-1; i++ {
		for j := 1; j < len(tree_map[0])-1; j++ {
			curr_scenic_score := d8p2_scenic_score(tree_map, i, j)
			if curr_scenic_score > scenic_score {
				scenic_score = curr_scenic_score
			}
		}
	}
	log.Printf("Highest scenic score: %d", scenic_score)
}

func init() {
	rootCmd.AddCommand(day8Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day8Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day8Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
