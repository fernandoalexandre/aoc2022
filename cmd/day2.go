/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// day2Cmd represents the day2 command
var day2Cmd = &cobra.Command{
	Use:   "day2",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day2 called")
		moves := d2p1_readfile("inputs/day2/input")
		d2p1(moves)
		d2p2(moves)
	},
}

var win_score int = 6

type move struct {
	other_player string
	me           string
}

func d2p1_readfile(file_name string) []move {
	readFile, err := os.Open(file_name)

	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []move

	for fileScanner.Scan() {
		moves := strings.Split(fileScanner.Text(), " ")
		switch moves[0] {
		case "A":
			moves[0] = "rock"
		case "B":
			moves[0] = "paper"
		case "C":
			moves[0] = "scissors"
		}

		switch moves[1] {
		case "X":
			moves[1] = "rock"
		case "Y":
			moves[1] = "paper"
		case "Z":
			moves[1] = "scissors"
		}

		fileLines = append(fileLines, move{moves[0], moves[1]})
	}

	readFile.Close()

	return fileLines
}

func d2p1_get_play_score(play string) int {
	switch play {
	case "rock":
		return 1
	case "paper":
		return 2
	case "scissors":
		return 3
	}
	return -1
}

func d2p1_calculate_score(mov move) int {
	var w_rules = make(map[string]string)

	w_rules["rock"] = "scissors"
	w_rules["paper"] = "rock"
	w_rules["scissors"] = "paper"

	if w_rules[mov.me] == mov.other_player {
		return win_score + d2p1_get_play_score(mov.me)
	} else if mov.me == mov.other_player {
		return (win_score / 2) + d2p1_get_play_score(mov.me)
	} else {
		return d2p1_get_play_score(mov.me)
	}
}

func d2p1(moves []move) {
	log.Println("Part 1")
	total_score := 0
	for _, curr_mov := range moves {
		total_score += d2p1_calculate_score(curr_mov)
	}
	log.Printf("Total score: %d", total_score)
}

func d2p2_calculate_score(mov move) int {
	var w_rules = make(map[string]string)
	var l_rules = make(map[string]string)

	w_rules["rock"] = "scissors"
	w_rules["paper"] = "rock"
	w_rules["scissors"] = "paper"

	l_rules["rock"] = "paper"
	l_rules["paper"] = "scissors"
	l_rules["scissors"] = "rock"

	switch mov.me {
	case "rock": // Need to lose "X"
		return d2p1_get_play_score(w_rules[mov.other_player])
	case "paper": // Need to draw "Y"
		return (win_score / 2) + d2p1_get_play_score(mov.other_player)
	case "scissors": // Need to win "Z"
		return win_score + d2p1_get_play_score(l_rules[mov.other_player])
	}
	return -1
}

func d2p2(moves []move) {
	log.Println("Part 2")
	total_score := 0
	for _, curr_mov := range moves {
		total_score += d2p2_calculate_score(curr_mov)
	}
	log.Printf("Total score: %d", total_score)
}

func init() {
	rootCmd.AddCommand(day2Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day2Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day2Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
