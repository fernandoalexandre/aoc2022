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

// day14Cmd represents the day14 command
var day14Cmd = &cobra.Command{
	Use:   "day14",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day14 called")

		test_content := d14_readfile("inputs/day14/input_example")
		content := d14_readfile("inputs/day14/input")

		d14p1(test_content)
		d14p1(content)

		f, err := os.OpenFile("testlogfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
		defer f.Close()

		log.SetOutput(f)

		//d14p2(test_content)
		content = d14_readfile("inputs/day14/input")
		d14p2(content)
	},
}

var SAND = '\u25E6'
var ORIGIN = '\u21E3'
var WALL = '\u2588'
var EMPTY = '\u2591'

func d14_readfile(file_name string) (content []string) {
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

func d14_print_map(vmap [][]rune) {
	separator := ""
	for i := 0; i < len(vmap[0]); i++ {
		separator += string('\u25AC')
	}
	log.Print(separator)
	for _, line := range vmap {
		log.Printf("%s", string(line))
	}
}

func d14_get_point_map_indexes(x, y, min_x, min_y int) (int, int) {
	// Convert max indexes (0,0) to (max_x-min_x, max_y-min_y) reference
	return x + min_x, y + min_y
}

func d14_get_map_indexes(x, y, min_x, min_y int) (int, int) {
	// Convert max indexes (max_x-min_x, max_y-min_y) to (0,0) reference
	return x - min_x, y - min_y
}

func d14p1_parse_map(content []string) (vmap [][]rune, min_x, min_y, max_x, max_y int) {
	min_x = 2147483647
	min_y = 0
	max_x = -2147483647
	max_y = -2147483647
	log.Println("First parse of the map (find min/max)")
	filled_points := map[string]int{}

	for _, line := range content {
		points := strings.Split(line, " -> ")
		start_point := &Point{-1, -1}
		dest_point := &Point{-1, -1}

		for _, point := range points {
			coords := strings.Split(point, ",")
			x, _ := strconv.Atoi(coords[0])
			y, _ := strconv.Atoi(coords[1])

			dest_point = start_point
			start_point = &Point{x, y}
			if x < min_x {
				min_x = x
			}

			if x > max_x {
				max_x = x
			}

			if y < min_y {
				min_y = y
			}

			if y > max_y {
				max_y = y
			}

			if dest_point.x != -1 {
				curr_min_y := start_point.y
				curr_max_y := dest_point.y
				if start_point.y > dest_point.y {
					curr_min_y = dest_point.y
					curr_max_y = start_point.y
				}

				curr_min_x := start_point.x
				curr_max_x := dest_point.x
				if start_point.x > dest_point.x {
					curr_min_x = dest_point.x
					curr_max_x = start_point.x
				}
				log.Printf("Going from (%d, %d) -> (%d, %d)", curr_min_x, curr_min_y, curr_max_x, curr_max_y)
				for curr_y := curr_min_y; curr_y <= curr_max_y; curr_y++ {
					for curr_x := curr_min_x; curr_x <= curr_max_x; curr_x++ {
						filled_points[d12_get_str(curr_x, curr_y)] = 0
					}
				}
			}
		}
	}
	log.Println("Second parse of the map (fill lines on map)")

	for y := 0; y < max_y-min_y+1; y++ {
		vmap_line := []rune{}
		for x := 0; x < max_x-min_x+1; x++ {
			new_x, new_y := d14_get_point_map_indexes(x, y, min_x, min_y)
			if _, ok := filled_points[d12_get_str(new_x, new_y)]; ok {
				vmap_line = append(vmap_line, WALL)
			} else {
				vmap_line = append(vmap_line, EMPTY)
			}
		}
		vmap = append(vmap, vmap_line)
	}

	return vmap, min_x, min_y, max_x, max_y
}

func d14p1(content []string) {
	log.Println("Part 1")

	source_x := 500
	source_y := 0

	vmap, min_x, min_y, max_x, max_y := d14p1_parse_map(content)

	new_x, new_y := d14_get_map_indexes(source_x, source_y, min_x, min_y)
	vmap[new_y][new_x] = ORIGIN

	d14_print_map(vmap)
	curr_sand := &Point{new_x, new_y + 1}
	has_finished := false
	stable_sands := 0
	for !has_finished { // For many sands
		for { // For this sand to stabilize
			bottom_slot := vmap[curr_sand.y+1][curr_sand.x]
			if curr_sand.y+1 == len(vmap) {
				has_finished = true
				break
			}
			if bottom_slot == EMPTY { // Can fall further
				curr_sand.y++
			} else { // Bottom is filled
				if curr_sand.x == 0 { // Finished as it's going out of bounds
					has_finished = true
					break
				}
				if vmap[curr_sand.y+1][curr_sand.x-1] == EMPTY {
					curr_sand = &Point{curr_sand.x - 1, curr_sand.y + 1}
				} else if curr_sand.x == len(vmap[0])-1 {
					has_finished = true
					break
				} else if vmap[curr_sand.y+1][curr_sand.x+1] == EMPTY {
					curr_sand = &Point{curr_sand.x + 1, curr_sand.y + 1}
				} else { // Can rest here
					vmap[curr_sand.y][curr_sand.x] = SAND
					stable_sands++
					curr_sand = &Point{new_x, new_y + 1}
					if vmap[new_x][new_y+1] != EMPTY {
						has_finished = true
					}
					break
				}
			}
		}
		d14_print_map(vmap)
	}

	log.Printf("Number of sands: %d", stable_sands)
	log.Printf("%d | %d | %d | %d", min_x, min_y, max_x, max_y)
}

func d14p2_parse_map(content []string) (vmap [][]rune, max_y int, x_offset int) {
	max_y = -2147483647
	log.Println("First parse of the map (find min/max)")
	filled_points := map[string]int{}

	for _, line := range content {
		points := strings.Split(line, " -> ")
		start_point := &Point{-1, -1}
		dest_point := &Point{-1, -1}

		for _, point := range points {
			coords := strings.Split(point, ",")
			x, _ := strconv.Atoi(coords[0])
			y, _ := strconv.Atoi(coords[1])

			dest_point = start_point
			start_point = &Point{x, y}

			if y > max_y {
				max_y = y
			}

			if dest_point.x != -1 {
				curr_min_y := start_point.y
				curr_max_y := dest_point.y
				if start_point.y > dest_point.y {
					curr_min_y = dest_point.y
					curr_max_y = start_point.y
				}

				curr_min_x := start_point.x
				curr_max_x := dest_point.x
				if start_point.x > dest_point.x {
					curr_min_x = dest_point.x
					curr_max_x = start_point.x
				}
				log.Printf("Going from (%d, %d) -> (%d, %d)", curr_min_x, curr_min_y, curr_max_x, curr_max_y)
				for curr_y := curr_min_y; curr_y <= curr_max_y; curr_y++ {
					for curr_x := curr_min_x; curr_x <= curr_max_x; curr_x++ {
						filled_points[d12_get_str(curr_x, curr_y)] = 0
					}
				}
			}
		}
	}
	log.Println("Second parse of the map (fill lines on map)")

	x_offset = 100

	for y := 0; y < max_y+3; y++ {
		vmap_line := []rune{}
		for x := -x_offset; x < x_offset+500; x++ {
			if y == max_y+2 {
				vmap_line = append(vmap_line, WALL)
			} else if _, ok := filled_points[d12_get_str(x+x_offset, y)]; ok {
				vmap_line = append(vmap_line, WALL)
			} else {
				vmap_line = append(vmap_line, EMPTY)
			}
		}
		vmap = append(vmap, vmap_line)
	}

	return vmap, max_y, x_offset
}

func d14p2(content []string) {
	log.Println("Part 2")

	source_x := 500
	source_y := 0

	vmap, _, _ := d14p2_parse_map(content)

	vmap[source_y][source_x] = ORIGIN

	d14_print_map(vmap)
	curr_sand := &Point{source_x, source_y}
	has_finished := false
	stable_sands := 0
	for !has_finished { // For many sands
		for { // For this sand to stabilize
			bottom_slot := vmap[curr_sand.y+1][curr_sand.x]
			if curr_sand.y+1 == len(vmap) {
				has_finished = true
				break
			}
			if bottom_slot == EMPTY { // Can fall further
				curr_sand.y++
			} else { // Bottom is filled
				if curr_sand.x == 0 { // Finished as it's going out of bounds
					has_finished = true
					break
				}
				if vmap[curr_sand.y+1][curr_sand.x-1] == EMPTY {
					curr_sand = &Point{curr_sand.x - 1, curr_sand.y + 1}
				} else if curr_sand.x == len(vmap[0])-1 {
					has_finished = true
					break
				} else if vmap[curr_sand.y+1][curr_sand.x+1] == EMPTY {
					curr_sand = &Point{curr_sand.x + 1, curr_sand.y + 1}
				} else { // Can rest here
					vmap[curr_sand.y][curr_sand.x] = SAND
					stable_sands++
					curr_sand = &Point{source_x, source_y}
					if vmap[source_y][source_x] != ORIGIN {
						has_finished = true
					}
					break
				}
			}
		}
	}
	d14_print_map(vmap)

	log.Printf("Number of sands: %d", stable_sands)
}

func init() {
	rootCmd.AddCommand(day14Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day14Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day14Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
