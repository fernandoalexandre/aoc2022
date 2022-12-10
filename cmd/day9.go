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

		test_commands := []string{"R 4",
			"U 4",
			"L 3",
			"D 1",
			"R 4",
			"D 1",
			"L 5",
			"R 2"}
		commands := d9_readfile("inputs/day9/input")
		rope := []Point{{0, 0}, {0, 0}}

		d9(test_commands, rope)
		rope = []Point{{0, 0}, {0, 0}}
		d9(commands, rope)

		test_commands_2 := []string{"R 5",
			"U 8",
			"L 8",
			"D 3",
			"R 17",
			"D 10",
			"L 25",
			"U 20"}
		rope = []Point{{0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}}
		d9(test_commands, rope)
		rope = []Point{{0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}}
		d9(test_commands_2, rope)
		rope = []Point{{0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}}
		d9(commands, rope)

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

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func d9_should_move(head, tail int) bool {
	return Abs(head-tail) > 1
}

func d9_move(rope *[]Point, tail_unique_positions map[Point]int, dir string) map[Point]int {
	switch dir {
	case "U":
		(*rope)[0].y += 1
	case "D":
		(*rope)[0].y -= 1
	case "L":
		(*rope)[0].x -= 1
	case "R":
		(*rope)[0].x += 1
	}

	for i := 1; i < len(*rope); i++ {
		if d9_should_move((*rope)[i-1].y, (*rope)[i].y) || d9_should_move((*rope)[i-1].x, (*rope)[i].x) {
			if (*rope)[i-1].x > (*rope)[i].x {
				(*rope)[i].x += 1
			} else if (*rope)[i-1].x < (*rope)[i].x {
				(*rope)[i].x -= 1
			}
			if (*rope)[i-1].y > (*rope)[i].y {
				(*rope)[i].y += 1
			} else if (*rope)[i-1].y < (*rope)[i].y {
				(*rope)[i].y -= 1
			}
		}
	}

	tail_unique_positions[Point{(*rope)[len(*rope)-1].x, (*rope)[len(*rope)-1].y}] = 0

	//log.Printf("Tail Position: %d | %v", len(tail_positions), tail_positions)
	return tail_unique_positions
}

func d9_print_map(last_pos_map map[Point]int) {
	for i := -15; i < 15; i++ {
		line := ""
		for j := -15; j < 15; j++ {
			_, found := last_pos_map[Point{i, j}]
			if i == 0 && j == 0 {
				line += "s"
			} else if !found {
				line += "."
			} else {
				line += "#"
			}
		}
		log.Println(line)
	}
}

func d9(commands []string, rope []Point) {
	fmt.Println("Part 1")

	tail_unique_positions := make(map[Point]int)
	tail_unique_positions[Point{0, 0}] = 0

	for _, cmd := range commands {
		re := regexp.MustCompile(`(.*) (\d+)`)
		p_cmd := re.FindStringSubmatch(cmd)

		moves, _ := strconv.Atoi(p_cmd[2])
		for i := 0; i < moves; i++ {
			tail_unique_positions = d9_move(&rope, tail_unique_positions, p_cmd[1])
			//log.Printf("rope = %v (%s)", rope, cmd)
			// Debug maps
			//point_map := make(map[Point]int)
			//for _, point := range rope {
			//	point_map[point] = 0
			//}
			//d9_print_point_map(point_map)
		}

	}

	// Debug final positions
	// log.Printf("Tail Positions: %v", tail_unique_positions)
	log.Printf("Tail Position: %d", len(tail_unique_positions))
	d9_print_map(tail_unique_positions)
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
