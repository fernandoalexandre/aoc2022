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

	"github.com/spf13/cobra"
)

// day20Cmd represents the day20 command
var day20Cmd = &cobra.Command{
	Use:   "day20",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day20 called")

		test_content := day20_readfile("inputs/day20/input_example")
		content := day20_readfile("inputs/day20/input")

		day20p1(test_content)
		day20p1(content)

		day20p2(content)

	},
}

func day20_readfile(file_name string) (content []int) {
	readFile, err := os.Open(file_name)

	if err != nil {
		log.Fatal(err)
	}

	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		val, _ := strconv.Atoi(fileScanner.Text())
		content = append(content, val)
	}
	return content
}

type CipherEntry struct {
	orig_pos, curr_pos, val int
	alt_val                 int64
}

func d20_print_state(state []*CipherEntry) {
	for _, entry := range state {
		fmt.Print(entry.alt_val)
		fmt.Print(" ")
	}
	fmt.Print("\n")

	for _, entry := range state {
		fmt.Print(entry.val)
		fmt.Print(" ")
	}
	fmt.Print("\n")

	for _, entry := range state {
		fmt.Print(entry.curr_pos)
		fmt.Print(" ")
	}
	fmt.Print("\n")
}

func insert(a []*CipherEntry, index int, value *CipherEntry) []*CipherEntry {
	if len(a) == index { // nil or empty slice or after last element
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...) // index < len(a)
	a[index] = value
	return a
}

func day20p1(commands []int) {
	log.Println("Part 1")

	to_process := []*CipherEntry{}
	state := []*CipherEntry{}
	for idx, cmd := range commands {
		entry := CipherEntry{idx, idx, cmd, int64(cmd)}
		to_process = append(to_process, &entry)
		state = append(state, &entry)
	}

	//d20_print_state(state)

	length := len(to_process)
	for i := 0; i < length; i++ {
		// Remove entry
		//log.Println("new")
		if to_process[i].val != 0 {
			//log.Printf("Curr %d", to_process[i].curr_pos)
			state = append(state[:to_process[i].curr_pos], state[to_process[i].curr_pos+1:]...)
			//d20_print_state(state)

			target_position := (to_process[i].curr_pos + to_process[i].val)
			if target_position < 0 {
				for target_position < 0 {
					target_position = length + target_position - 1
				}
			} else if target_position == 0 && to_process[i].val < 0 {
				target_position = length - 1
			} else if target_position > length {
				for target_position > length {
					target_position -= length
				}
				target_position = (target_position + 1) % length
			} else if target_position == length {
				target_position = 1
			}

			to_process[i].curr_pos = target_position

			state = insert(state, target_position, to_process[i])

			//log.Printf("%d", target_position)
			for j := 0; j < len(to_process); j++ {
				state[j].curr_pos = j
			}
			//d20_print_state(state)
		}
	}

	final_idx := -1
	for i := 0; i < len(state); i++ {
		if state[i].val == 0 {
			final_idx = i
			break
		}
	}

	result := 0
	//d20_print_state(state)
	i := []int{1000, 2000, 3000}
	for _, idx := range i {
		curr_val := state[(final_idx+idx)%len(state)].val
		log.Printf("Sum: %d | %d", curr_val, (final_idx+idx)%len(state))
		result += curr_val
	}

	log.Printf("Result: %d", result)
}

func day20p2(commands []int) {
	log.Println("Part 2")

	to_process := []*CipherEntry{}
	state := []*CipherEntry{}
	factor := 811589153 % (len(commands) - 1)
	for idx, cmd := range commands {
		entry := CipherEntry{idx, idx, cmd, int64(cmd * factor)}
		to_process = append(to_process, &entry)
		state = append(state, &entry)
	}

	//d20_print_state(state)

	for k := 0; k < 10; k++ {
		for i := 0; i < len(to_process); i++ {
			// Remove entry
			//log.Println("new")
			if to_process[i].alt_val != 0 {
				//log.Printf("Curr %d", to_process[i].curr_pos)
				//d20_print_state(state)
				//log.Println(to_process[i].curr_pos)

				state = append(state[:to_process[i].curr_pos], state[to_process[i].curr_pos+1:]...)
				//d20_print_state(state)

				length := int64(len(to_process))
				target_position := (int64(to_process[i].curr_pos) + to_process[i].alt_val)
				if target_position < 0 {
					for target_position < 0 {
						target_position = length + target_position - 1
					}
				} else if target_position > length {
					for target_position > length {
						target_position -= length - 1
					}
				}

				if target_position == 0 && to_process[i].alt_val < 0 {
					target_position = length - 1
				} else if target_position == length {
					target_position = 1
				}

				//log.Println(target_position)

				state = insert(state, int(target_position), to_process[i])

				//log.Printf("%d", target_position)
				for j := 0; j < len(to_process); j++ {
					state[j].curr_pos = j
				}
				//d20_print_state(state)
			}
		}
		//d20_print_state(state)
	}

	final_idx := -1
	for i := 0; i < len(state); i++ {
		if state[i].alt_val == 0 {
			final_idx = i
			break
		}
	}

	var result int64 = 0
	//d20_print_state(state)
	i := []int{1000, 2000, 3000}
	for _, idx := range i {
		var curr_val int64 = int64(state[(final_idx+idx)%len(state)].val) * 811589153
		log.Printf("Sum: %d | %d", curr_val, (final_idx+idx)%len(state))
		result += curr_val
	}

	//d20_print_state(state)

	log.Printf("Result: %d", result)
}

func init() {
	rootCmd.AddCommand(day20Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day20Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day20Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
