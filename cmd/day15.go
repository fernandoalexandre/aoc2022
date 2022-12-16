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

// day15Cmd represents the day15 command
var day15Cmd = &cobra.Command{
	Use:   "day15",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day15 called")

		test_content, occupied_test := d15_readfile("inputs/day15/input_example")
		d15p1(test_content, occupied_test, 10)
		d15p1v2(test_content, occupied_test, 10)

		content, occupied := d15_readfile("inputs/day15/input")
		d15p1(content, occupied, 2000000)
		d15p1v2(content, occupied, 2000000)

		d15p2(test_content, occupied_test, 0, 20)
		d15p2(content, occupied, 0, 4000000)
	},
}

type Sensor struct {
	sensor_pt      Point
	closest_beacon Point
	distance       int
}

func (a Sensor) intersect_area(pt Point) bool {
	return a.distance >= d15_mdistance(a.sensor_pt, pt)
}

func d15_mdistance(first, second Point) int {
	return Abs(first.x-second.x) + Abs(first.y-second.y)
}

func (a Sensor) get_xs_for_fixed_y(y int) (left, right int) {
	return a.sensor_pt.x - (a.distance - Abs(y-a.sensor_pt.y)), a.sensor_pt.x + (a.distance - Abs(y-a.sensor_pt.y))
}

func d15_readfile(file_name string) (content []*Sensor, occupied map[string]Point) {
	readFile, err := os.Open(file_name)

	if err != nil {
		log.Fatal(err)
	}

	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	occupied = map[string]Point{}

	for fileScanner.Scan() {
		re := regexp.MustCompile(`Sensor at x=(-?[0-9]\d*), y=(-?[0-9]\d*): closest beacon is at x=(-?[0-9]\d*), y=(-?[0-9]\d*)`)
		matches := re.FindStringSubmatch(fileScanner.Text())

		s_x, _ := strconv.Atoi(matches[1])
		s_y, _ := strconv.Atoi(matches[2])
		b_x, _ := strconv.Atoi(matches[3])
		b_y, _ := strconv.Atoi(matches[4])
		sensor := Point{s_x, s_y}
		beacon := Point{b_x, b_y}

		occupied[d12_get_str(s_x, s_y)] = sensor
		occupied[d12_get_str(b_x, b_y)] = beacon
		content = append(content, &Sensor{sensor, beacon, d15_mdistance(sensor, beacon)})
	}
	return content, occupied
}

func d15p1(sensor_list []*Sensor, occupied map[string]Point, y_idx int) {
	log.Println("Part 1 -- Semi brute force")

	intersections := map[string]bool{}

	for _, sensor := range sensor_list {
		max_distance := d15_mdistance(sensor.sensor_pt, Point{sensor.sensor_pt.x, y_idx})
		if max_distance <= sensor.distance { //  It's within the range counting vertically (most distance)
			_, ok := occupied[d12_get_str(sensor.sensor_pt.x, y_idx)]
			if !ok {
				intersections[d12_get_str(sensor.sensor_pt.x, y_idx)] = true
			}

			for x := 1; x <= sensor.sensor_pt.x+max_distance*4; x++ {
				// Start in the 'middle' (X = sensor.x + 1) and expands to both sides evenly
				point := Point{sensor.sensor_pt.x + x, y_idx}
				if sensor.intersect_area(point) {
					_, ok := occupied[d12_get_str(point.x, point.y)]
					if !ok {
						intersections[d12_get_str(point.x, point.y)] = true
					}
					_, ok = occupied[d12_get_str(sensor.sensor_pt.x-x, point.y)]
					if !ok {
						intersections[d12_get_str(sensor.sensor_pt.x-x, point.y)] = true
					}
				} else {
					break
				}
			}
		}
	}

	log.Printf("Number of intersections: %d", len(intersections))
}

func d15p1v2(sensor_list []*Sensor, occupied map[string]Point, y_idx int) {
	log.Println("Part 1v2, mathzzzz")
	// This approach was done only to validate the get_xs_for_fixed_y() function for part 2, doesn't run faster

	intersections := map[string]bool{}

	for _, sensor := range sensor_list {
		max_distance := d15_mdistance(sensor.sensor_pt, Point{sensor.sensor_pt.x, y_idx})
		if max_distance <= sensor.distance { //  It's within the range counting vertically (most distance)
			left, right := sensor.get_xs_for_fixed_y(y_idx)
			log.Printf("%v | Left: %d | Right %d", sensor, left, right)

			for x := left; x <= right; x++ {
				_, ok := occupied[d12_get_str(x, y_idx)]
				if !ok {
					intersections[d12_get_str(x, y_idx)] = true
				}

			}
		}
	}

	log.Printf("Number of intersections: %d", len(intersections))
}

func d15p2(sensor_list []*Sensor, occupied map[string]Point, MIN, MAX int) {
	log.Println("Part 2")

	for y := MIN; y <= MAX; y++ {
		if y%5000 == 0 {
			log.Printf("Iteration %d/%d", y, MAX)
		}
		found := false
		for _, sensor := range sensor_list {
			left, right := sensor.get_xs_for_fixed_y(y)

			if left > MIN {
				if left < MAX {
					for _, o_sensor := range sensor_list {
						if o_sensor == sensor {
							continue
						}

						if d15_mdistance(o_sensor.sensor_pt, Point{left - 1, y}) <= o_sensor.distance {
							found = true
							break
						}
					}
					if !found {
						log.Printf("Hooray! @ %v", Point{left - 1, y})
						log.Printf("Tuning: %d", (left-1)*MAX+y)
						return
					} else {
						found = false
					}
				} else {
					continue
				}
			}

			if right > MIN {
				if right < MAX {
					for _, o_sensor := range sensor_list {
						if o_sensor == sensor {
							continue
						}

						if d15_mdistance(o_sensor.sensor_pt, Point{right + 1, y}) <= o_sensor.distance {
							found = true
							break
						}
					}
					if !found {
						log.Printf("Hooray! @ %v", Point{right + 1, y})
						log.Printf("Tuning: %d", (right+1)*MAX+y)
						return
					} else {
						found = false
					}
				} else {
					continue
				}
			}
		}

		if found {
			break
		}
	}
}

func init() {
	rootCmd.AddCommand(day15Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day15Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day15Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
