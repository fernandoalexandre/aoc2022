/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"

	"github.com/spf13/cobra"
)

// day9Cmd represents the day9 command
var day9Cmd = &cobra.Command{
	Use:   "day9",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day9 called")

		test_commands := []string{}
		commands := d9_readfile("inputs/day9/input")
		d9p1(test_commands)
		d9p1(commands)

		d9p2(test_commands)
		d9p2(commands)

	},
}

type Point struct {
	x int
	y int
}

func d9_readfile(file_name string) (content []string) {
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

func d9p1_should_move(head, tail *Point) bool {
	return false
}

func d9p1_move(head, tail *Point, head_positions, tail_positions *[]Point, dir string, moves int) {
	old_head := Point{head.x, head.y}
	head_pos := append(*head_positions, Point{head.x, head.y})
	head_positions = &head_pos
	switch dir {
	case "U":
		head = &Point{head.x, head.y + 1}
	case "D":
		head = &Point{head.x, head.y - 1}
	case "L":
		head = &Point{head.x - 1, head.y}
	case "R":
		head = &Point{head.x + 1, head.y}
	}

	if d9p1_should_move(head, tail) {
		tail_pos := append(*tail_positions, Point{tail.x, tail.y})
		tail_positions = &tail_pos
		tail.x = old_head.x
		tail.y = old_head.y
	}
}

func d9p1(commands []string) {
	fmt.Println("Part 1")
	head := Point{0, 0}
	tail := Point{0, 0}
	head_positions := []Point{}
	tail_positions := []Point{}

	for _, cmd := range commands {
		re := regexp.MustCompile(`(.*) (\d+)`)
		p_cmd := re.FindStringSubmatch(cmd)

		moves, _ := strconv.Atoi(p_cmd[2])
		d9p1_move(&head, &tail, &head_positions, &tail_positions, p_cmd[1], moves)
	}

	log.Printf("Tail Position: %d | %v", len(tail_positions), tail_positions)
}

func d9p2(commands []string) {
	fmt.Println("Part 1")
}

func init() {
	rootCmd.AddCommand(day9Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day9Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day9Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
