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

// day5Cmd represents the day5 command
var day5Cmd = &cobra.Command{
	Use:   "day5",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day5 called")
		state, commands := d5_readfile("inputs/day5/input")

		d5p1(state, commands)
		d5p2(state, commands)
	},
}

// Common Functions
func d5_parse_state(state []string) (stacks []Stack) {
	log.Println("Parsing stack state...")

	last_stack_idx := len(state) - 1
	last_stack := state[last_stack_idx]
	num_stacks := 0
	// Find out how many stacks
	for i := 1; i <= len(last_stack); i += 4 {
		num_stacks++
	}

	stacks = make([]Stack, num_stacks)

	// Foreach line (inverted)
	for i := last_stack_idx - 1; i >= 0; i-- {
		// For each 3 characters
		k := 0
		for j := 1; j <= len(state[i]); j += 4 {
			if state[i][j] != ' ' {
				stacks[k].Push(string(state[i][j]))
			}
			k++
		}
	}

	return stacks
}

func d5_readfile(file_name string) (state, commands []string) {
	readFile, err := os.Open(file_name)

	if err != nil {
		log.Fatal(err)
	}

	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	is_map := true

	for fileScanner.Scan() {
		if fileScanner.Text() == "" {
			is_map = false
			continue
		}

		if is_map {
			state = append(state, fileScanner.Text())
		} else {
			commands = append(commands, fileScanner.Text())
		}
	}
	return state, commands
}

func d5_get_command_matches(cmd string) (n_slots, from, to int) {
	re := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
	p_cmd := re.FindStringSubmatch(cmd)

	n_slots, _ = strconv.Atoi(p_cmd[1])
	from, _ = strconv.Atoi(p_cmd[2])
	to, _ = strconv.Atoi(p_cmd[3])

	return n_slots, from, to
}

func d5_print_top_items(stacks []Stack) {
	log.Print("Top Containers: ")

	output := ""
	for _, stack := range stacks {
		val, has_val := stack.Pop()
		if !has_val {
			output += " "
		} else {
			output += val
		}
	}
	log.Println(output)
}

// Part 1
func d5p1_run_command(state *[]Stack, command string) {
	n_slots, from, to := d5_get_command_matches(command)

	for i := 0; i < n_slots; i++ {
		if (*state)[from-1].IsEmpty() {
			break
		}
		val, _ := (*state)[from-1].Pop()

		(*state)[to-1].Push(val)
	}
}

func d5p1(state []string, commands []string) {
	log.Println("Part 1")

	stacks := d5_parse_state(state)
	log.Printf("%v", stacks)

	for _, command := range commands {
		d5p1_run_command(&stacks, command)
	}
	log.Printf("%v", stacks)

	d5_print_top_items(stacks)
}

// Part 2
func d5p2_run_command(state *[]Stack, command string) {
	n_slots, from, to := d5_get_command_matches(command)

	val, _ := (*state)[from-1].Multipop(n_slots)

	(*state)[to-1].Multipush(val)
}

func d5p2(state []string, commands []string) {
	log.Println("Part 2")
	stacks := d5_parse_state(state)

	log.Printf("%v", stacks)
	for _, command := range commands {
		d5p2_run_command(&stacks, command)
	}
	log.Printf("%v", stacks)

	d5_print_top_items(stacks)

}

func init() {
	rootCmd.AddCommand(day5Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day5Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day5Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// Borrowed from https://www.educative.io/answers/how-to-implement-a-stack-in-golang
type Stack []string

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(str string) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

// Adapted for AOC
// Remove and return N elements from the top of stack. Return false if stack is empty.
func (s *Stack) Multipop(n_entries int) ([]string, bool) {
	if s.IsEmpty() {
		return []string{}, false
	} else {
		if n_entries > len(*s) {
			n_entries = len(*s)
		}
		index := len(*s) - n_entries // Get the index of the new top most element.
		elements := (*s)[index:]     // Index into the slice and obtain the elements.
		*s = (*s)[:index]            // Remove it from the stack by slicing it off.
		return elements, true
	}
}

// Push a new values onto the stack maintaining the list's order
func (s *Stack) Multipush(str []string) {
	for _, val := range str {
		*s = append(*s, val)
	}
}
