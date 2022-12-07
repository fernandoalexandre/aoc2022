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

// day7Cmd represents the day7 command
var day7Cmd = &cobra.Command{
	Use:   "day7",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day7 called")
		test_commands := []string{"$ cd /",
			"$ ls",
			"dir a",
			"14848514 b.txt",
			"8504156 c.dat",
			"dir d",
			"$ cd a",
			"$ ls",
			"dir e",
			"29116 f",
			"2557 g",
			"62596 h.lst",
			"$ cd e",
			"$ ls",
			"584 i",
			"$ cd ..",
			"$ cd ..",
			"$ cd d",
			"$ ls",
			"4060174 j",
			"8033020 d.log",
			"5626152 d.ext",
			"7214296 k"}
		log.Println("Test Part 1")
		d7p1(test_commands)
		log.Println("Actual Part 1")
		d7p1(d7_readfile("inputs/day7/input"))

		log.Println("Test Part 2")
		d7p2(test_commands)
		log.Println("Actual Part 2")
		d7p2(d7_readfile("inputs/day7/input"))
	},
}

func d7_readfile(file_name string) (content []string) {
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

func print_fs(fs *Node, ident int) {
	prefix := ""
	for i := 0; i < ident; i++ {
		prefix += "  "
	}

	log.Printf("%s- %s", prefix, fs.str())

	for _, child := range fs.children {
		print_fs(child, ident+1)
	}
}

func d7_create_fs(fs *Node, commands []string) {
	var curr_node *Node = fs
	for _, command := range commands[1:] {
		if command[0] == '$' { // It's a command
			if command[2:4] == "cd" {
				if command[5:] == ".." {
					curr_node = curr_node.parent
				} else {
					n := Node{command[5:], 0, nil, []*Node{}}

					curr_node.add_child(&n)
					curr_node = &n
				}
			} else if command[2:4] == "ls" {
				continue
			} else {
				log.Printf("Unknown command %s", command[2:4])
			}
		} else {
			re := regexp.MustCompile(`(\d+) (.*)`)
			p_cmd := re.FindStringSubmatch(command)

			if len(p_cmd) > 1 { // Only care about file sizes right now
				size, _ := strconv.Atoi(p_cmd[1])
				curr_node.add_child(&Node{p_cmd[2], size, nil, []*Node{}})
			}
		}
	}

	fs.calc_size()

	print_fs(fs, 0)
}

func d7p1_recursive_calc(fs *Node, curr_size *int) {
	if fs.is_dir() {
		if fs.size < 100000 {
			*curr_size += fs.size
		}
		for _, child := range fs.children {
			d7p1_recursive_calc(child, curr_size)
		}
	}
}

func d7p1(commands []string) {
	log.Printf("Part 1")

	var fs Node = Node{"/", 0, nil, []*Node{}}

	d7_create_fs(&fs, commands)

	total := 0
	d7p1_recursive_calc(&fs, &total)
	log.Printf("Sum of sizes lower or equal than 100000: %d", total)
}

func d7p2_check_folder_to_delete(fs *Node, delete_threshold int) int {
	if fs.is_dir() {
		if fs.size >= delete_threshold {
			current_min := fs.size
			for _, child := range fs.children {
				child_min_size := d7p2_check_folder_to_delete(child, delete_threshold)
				if child_min_size > 0 && child_min_size < current_min {
					current_min = child_min_size
				}
			}
			return current_min
		}
	}
	return -1
}

func d7p2(commands []string) {
	log.Printf("Part 2")

	var fs Node = Node{"/", 0, nil, []*Node{}}

	d7_create_fs(&fs, commands)

	// Get the size we need to delete before reaching the 30000000 threshold
	delete_threshold := 30000000 - (70000000 - fs.size)

	log.Printf("Delete Threshold: %d", delete_threshold)

	min_delete := d7p2_check_folder_to_delete(&fs, delete_threshold)

	log.Printf("Minimum size to delete: %d", min_delete)
}

func init() {
	rootCmd.AddCommand(day7Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day7Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day7Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type Node struct {
	name     string
	size     int
	parent   *Node
	children []*Node
}

func (a *Node) add_child(child *Node) {
	child.parent = a
	a.children = append(a.children, child)
}

func (a *Node) calc_size() {
	if len(a.children) > 0 {
		a.size = 0
		for _, child := range a.children {
			child.calc_size()
			a.size += child.size
		}
	}
}

func (a *Node) is_dir() bool {
	return len(a.children) > 0
}

func (a *Node) str() string {
	suffix := ""
	if len(a.children) > 0 {
		suffix = fmt.Sprintf("(dir, size=%d)", a.size)
	} else {
		suffix = fmt.Sprintf("(file, size=%d)", a.size)
	}

	return fmt.Sprintf("%s %s", a.name, suffix)
}
