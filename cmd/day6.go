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

// day6Cmd represents the day6 command
var day6Cmd = &cobra.Command{
	Use:   "day6",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day6 called")
		content := d6_readfile("inputs/day6/input")

		log.Print("Part 1")
		log.Printf("Start tests")

		d6("mjqjpqmgbljsphdztnvjfqwrcgsmlb", 4)
		d6("bvwbjplbgvbhsrlpgdmjqwftvncz", 4)
		d6("nppdvjthqldpwncqszvftbrmjlhg", 4)
		d6("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 4)
		d6("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 4)

		log.Printf("Real deal")
		d6(content[0], 4)

		log.Print("Part 2")
		log.Printf("Start tests")
		d6("mjqjpqmgbljsphdztnvjfqwrcgsmlb", 14)
		d6("bvwbjplbgvbhsrlpgdmjqwftvncz", 14)
		d6("nppdvjthqldpwncqszvftbrmjlhg", 14)
		d6("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 14)
		d6("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 14)

		log.Printf("Real deal")
		d6(content[0], 14)
	},
}

func d6_readfile(file_name string) (signals []string) {
	readFile, err := os.Open(file_name)

	if err != nil {
		log.Fatal(err)
	}

	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		signals = append(signals, fileScanner.Text())
	}
	return signals
}

func d6(signal string, length int) {
	for i := length - 1; i < len(signal); i++ {
		curr_map := make(map[byte]int)
		has_equal := false
		for j := length - 1; j >= 0; j-- {
			_, ok := curr_map[signal[i-j]]
			if ok {
				has_equal = true
				break
			} else {
				curr_map[signal[i-j]] = 0
			}
		}

		if !has_equal {
			log.Printf("Found sequence %s, in index %d", signal[i-length:i+1], i+1)
			break
		}
	}
}

func init() {
	rootCmd.AddCommand(day6Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day6Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day6Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
