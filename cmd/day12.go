/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"

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
		d12p1(actual_map)

		d12p2(test_map)
		d12p2(actual_map)
	},
}

type MapNode struct {
	token       rune
	x, y, value int
	links       map[string]*MapNode
}

func (a *MapNode) str() string {
	return fmt.Sprintf("%d,%d", a.x, a.y)
}

func d12_get_str(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

func (a *MapNode) d12_can_move(target MapNode) bool {
	// Goin' up
	if a.token <= target.token {
		return int(target.token)-int(a.token) <= 1
	}
	// Goin' down
	return true
}

func create_map_node(slot rune, x, y int) (node *MapNode, is_start, is_dest bool) {
	node = &MapNode{slot, x, y, 2147483647, map[string]*MapNode{}}
	if slot == 'S' {
		node.token = 'a'
		is_dest = false
		is_start = true
	} else if slot == 'E' {
		node.token = 'z'
		is_dest = true
		is_start = false
	} else {
		is_dest = false
		is_start = false
	}
	return node, is_start, is_dest
}

func d12_parse_map(input_map []string) (output_map map[string]*MapNode, start_node *MapNode, dest_node *MapNode, start_node_list []*MapNode) {
	output_map = make(map[string]*MapNode)
	start_node_list = []*MapNode{}
	for y, line := range input_map {
		for x, slot := range line {
			val, found := output_map[d12_get_str(x, y)]
			is_start := false
			is_dest := false
			if !found {
				val, is_start, is_dest = create_map_node(slot, x, y)

				if is_start {
					start_node = val
				} else if is_dest {
					dest_node = val
				}

				output_map[val.str()] = val
			}

			if val.token == 'a' {
				start_node_list = append(start_node_list, val)
			}

			// x y+1
			if y < len(input_map)-1 {
				//log.Printf("Adding (%s)", d12_get_str(x, y+1))
				b_val, found := output_map[d12_get_str(x, y+1)]
				if !found {
					b_val, is_start, is_dest = create_map_node(rune(input_map[y+1][x]), x, y+1)
					if is_start {
						start_node = b_val
					} else if is_dest {
						dest_node = b_val
					}
					output_map[b_val.str()] = b_val
				}

				if val.d12_can_move(*b_val) {
					val.links[b_val.str()] = b_val
				}
				if b_val.d12_can_move(*val) {
					b_val.links[val.str()] = val
				}
			}

			// x+1 y
			if x < len(line)-1 {
				r_val, found := output_map[d12_get_str(x+1, y)]
				if !found {
					r_val, is_start, is_dest = create_map_node(rune(input_map[y][x+1]), x+1, y)
					output_map[r_val.str()] = r_val
					if is_start {
						start_node = r_val
					} else if is_dest {
						dest_node = r_val
					}
				}

				if val.d12_can_move(*r_val) {
					val.links[r_val.str()] = r_val
				}
				if r_val.d12_can_move(*val) {
					r_val.links[val.str()] = val
				}
			}
		}
	}

	return output_map, start_node, dest_node, start_node_list
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

func find_shortest_routes(output_map map[string]*MapNode, start_node *MapNode, dest_node *MapNode) int {
	visited := make(map[string]bool)

	heap := &Heap{}

	start_node.value = 0
	heap.Push(start_node)

	for heap.Size() > 0 {
		curr_vertex := heap.Pop()

		if visited[curr_vertex.str()] {
			continue
		}

		visited[curr_vertex.str()] = true

		for _, neighbor := range curr_vertex.links {
			if !visited[neighbor.str()] {
				new_cost := curr_vertex.value + 1
				heap.Push(neighbor)

				if new_cost < neighbor.value {
					neighbor.value = new_cost
				}
			}
		}
	}
	//log.Printf("%v | %v", start_node.str(), start_node.links)
	//log.Printf("%v | %v | %v", dest_node.str(), dest_node.value, dest_node.links)

	return dest_node.value
}

func d12p1(raw_map []string) {
	log.Println("Part 1")
	output_map, start_node, dest_node, _ := d12_parse_map(raw_map)
	shortest_route := find_shortest_routes(output_map, start_node, dest_node)

	log.Printf("Shortest Route: %d", shortest_route)
}

func d12p2(raw_map []string) {
	log.Println("Part 2")
	output_map, _, dest_node, start_node_list := d12_parse_map(raw_map)
	curr_min := 2147483647
	for _, start_node := range start_node_list {
		shortest_route := find_shortest_routes(output_map, start_node, dest_node)

		if shortest_route < curr_min {
			curr_min = shortest_route
		}
	}

	log.Printf("Shortest Route: %d", curr_min)
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

// Min heap definition
// Adapted from https://go-recipes.dev/dijkstras-algorithm-in-go-e1129b2f5c9e

type Heap struct {
	elements []*MapNode
	mutex    sync.RWMutex
}

func (h *Heap) Size() int {
	h.mutex.RLock()
	defer h.mutex.RUnlock()
	return len(h.elements)
}

// push an element to the heap, re-arrange the heap
func (h *Heap) Push(element *MapNode) {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	h.elements = append(h.elements, element)
	i := len(h.elements) - 1
	for ; h.elements[i].value < h.elements[parent(i)].value; i = parent(i) {
		h.swap(i, parent(i))
	}
}

// pop the top of the heap, which is the min value
func (h *Heap) Pop() (i *MapNode) {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	i = h.elements[0]
	h.elements[0] = h.elements[len(h.elements)-1]
	h.elements = h.elements[:len(h.elements)-1]
	h.rearrange(0)
	return
}

// rearrange the heap
func (h *Heap) rearrange(i int) {
	smallest := i
	left, right, size := leftChild(i), rightChild(i), len(h.elements)
	if left < size && h.elements[left].value < h.elements[smallest].value {
		smallest = left
	}
	if right < size && h.elements[right].value < h.elements[smallest].value {
		smallest = right
	}
	if smallest != i {
		h.swap(i, smallest)
		h.rearrange(smallest)
	}
}

func (h *Heap) swap(i, j int) {
	h.elements[i], h.elements[j] = h.elements[j], h.elements[i]
}

func parent(i int) int {
	return (i - 1) / 2
}

func leftChild(i int) int {
	return 2*i + 1
}

func rightChild(i int) int {
	return 2*i + 2
}
