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

// day10Cmd represents the day10 command
var day10Cmd = &cobra.Command{
	Use:   "day10",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day10 called")

		test_content := d10_readfile("inputs/day10/input_example")
		content := d10_readfile("inputs/day10/input")
		log.Println("Test Case")
		d10(test_content)
		log.Println("Actual Exercise")
		d10(content)
	},
}

func d10_readfile(file_name string) (content []string) {
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

func d10(commands []string) {
	log.Println("Part 1")

	x := 1
	value_list := []int{}
	for _, command := range commands {
		p_cmd := strings.Split(command, " ")
		switch p_cmd[0] {
		case "noop":
			value_list = append(value_list, x)
		case "addx":
			value_list = append(value_list, x)
			val, _ := strconv.Atoi(p_cmd[1])
			x += val
			value_list = append(value_list, x)
		}
	}

	log.Printf("20: %d | 60: %d | 100: %d | 140: %d | 180: %d | 220: %d ", value_list[19], value_list[59], value_list[99], value_list[139], value_list[179], value_list[219])
	log.Printf("Strength: %d", value_list[19]*20+value_list[59]*60+value_list[99]*100+value_list[139]*140+value_list[179]*180+value_list[219]*220)

	log.Println("Part 2")
	current_strength := value_list[0]
	for i := 0; i < 6; i++ {
		line := ""
		for j := 0; j < 40; j++ {
			if j-1 <= current_strength && current_strength <= j+1 {
				line += "#"
			} else {
				line += "."
			}
			current_strength = value_list[i*40+j]
		}
		log.Println(line)
	}
}

func init() {
	rootCmd.AddCommand(day10Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day10Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day10Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
