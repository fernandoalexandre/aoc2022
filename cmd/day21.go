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
	"strings"

	"github.com/apaxa-go/eval"
	"github.com/spf13/cobra"
)

// day21Cmd represents the day21 command
var day21Cmd = &cobra.Command{
	Use:   "day21",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day21 called")

		test_exprs, test_vals := day21_readfile("inputs/day21/input_example")
		exprs, vals := day21_readfile("inputs/day21/input")

		day21p1(test_exprs, test_vals)
		day21p1(exprs, vals)

		//day21p2(test_exprs, test_vals)
		//day21p2(exprs, vals)

		day21p2_hardcoded()
	},
}

func day21_readfile(file_name string) (exprs map[string]string, vals map[string]string) {
	readFile, err := os.Open(file_name)

	if err != nil {
		log.Fatal(err)
	}

	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	re := regexp.MustCompile(`(.+): (.+)`)
	exprs, vals = map[string]string{}, map[string]string{}
	for fileScanner.Scan() {
		monkeh := re.FindStringSubmatch(fileScanner.Text())
		_, err := strconv.Atoi(monkeh[2])
		if err != nil {
			exprs[monkeh[1]] = monkeh[2]
		} else {
			vals[monkeh[1]] = monkeh[2]
		}
	}
	return exprs, vals
}

func day21p1(exprs map[string]string, vals map[string]string) {
	log.Println("Part 1")

	expression := exprs["root"]
	start_expression := ""
	for start_expression != expression {
		start_expression = expression
		for k, v := range exprs {
			if k != "root" {
				expression = strings.ReplaceAll(expression, k, fmt.Sprintf("(%s)", v))
			}
		}
	}

	for k, v := range vals {
		expression = strings.ReplaceAll(expression, k, v)
	}

	log.Println(expression)

	expr, err := eval.ParseString(expression, "")
	if err != nil {
		return
	}
	r, err := expr.EvalToInterface(nil)
	if err != nil {
		return
	}
	log.Printf("%v %T", r, r) // "3 int8"
}

func day21p2(exprs map[string]string, vals map[string]string) {
	log.Println("Part 2")

	expression := strings.Split(exprs["root"], "+")
	for j := 0; j < len(expression); j++ {
		start_expression := ""
		for start_expression != expression[j] {
			start_expression = expression[j]
			for k, v := range exprs {
				if k != "root" {
					expression[j] = strings.ReplaceAll(expression[j], k, fmt.Sprintf("(%s)", v))
				}
			}
		}
	}

	humn_idx := 0
	clear_idx := 1
	if strings.Contains(expression[1], "humn") {
		clear_idx = 0
		humn_idx = 1
	}
	for j := 0; j < len(expression); j++ {
		for k, v := range vals {
			if k != "humn" {
				expression[j] = strings.ReplaceAll(expression[j], k, v)
			}
		}
	}

	log.Println(expression[clear_idx])
	expr, err := eval.ParseString(expression[clear_idx], "")
	if err != nil {
		return
	}
	r, err := expr.EvalToInterface(nil)
	if err != nil {
		return
	}

	log.Printf("%f %T", r, r) // "3 int8"

	expression[humn_idx] = fmt.Sprintf("%.0f == %s", r, expression[humn_idx])
	log.Println(expression[humn_idx])
	humn := 3110000
	curr_expr := strings.ReplaceAll(expression[humn_idx], "humn", fmt.Sprint(humn))
	expr, err = eval.ParseString(curr_expr, "")
	if err != nil {
		return
	}
	r, err = expr.EvalToInterface(nil)
	if err != nil {
		return
	}
	for !r.(bool) {
		humn++
		curr_expr = strings.ReplaceAll(expression[humn_idx], "humn", fmt.Sprint(humn))
		expr, err = eval.ParseString(curr_expr, "")
		if err != nil {
			return
		}
		r, err = expr.EvalToInterface(nil)
		if err != nil {
			return
		}

		if humn%500 == 0 {
			log.Println(humn)
		}
	}
	log.Printf("%v %T", humn, r) // "3 int8"
}

func day21p2_hardcoded() {
	var lol bool = false
	var humn int64 = -1
	for !lol {
		humn++
		lol = 52716091087786 == ((((3 * (((((4 * 11) * 12) - (16 + (((3 * 3) * 2) * (2 + 5)))) * (((((2 * (((((15 + ((((4 + 5) - 2) * (1 + 15)) / 8)) * 2) / 2) + (2 * (12 + 2))) + (2 * 13))) + (3 * ((12 + 1) * ((19 - 2) + 12)))) + ((((((3 * 2) + ((5 + (6 + 17)) + 3)) - ((2 * 4) + 4)) * 2) * ((2 + 14) + (6 * (3 * 3)))) / 5)) * ((2 * (((((2 * 3) + (5 * 3)) - (2 + 5)) * (1 + (3 * 11))) + (6 * (6 + (1 + (6 * 3)))))) + (((((2 * ((9 * 3) + 10)) / 2) + (8 + 3)) + ((2 + (2 * 16)) - 9)) * (3 + 4)))) + ((((((3 * 4) + (12 + ((20 + (13 * 2)) + ((1 + (((3 + 10) * 2) * 2)) * 2)))) + (((20 * 2) + (4 + 3)) * 2)) + (3 * (5 * 5))) + ((((2 + 5) + ((3 * 3) * 3)) + 13) * 3)) * (17 * (2 * 9))))) - ((((4 * (12 + (1 + (2 * (5 * 3))))) + (((3 * (2 * 6)) + ((4 * 3) + 8)) + 11)) * 7) * (3 * ((((3 * 2) + 1) * (((((2 * 5) * 9) + (((2 + 5) + (((((2 * (((16 * 2) + (19 + (((2 * 3) + 1) * 4))) * 2)) / 4) * 2) / 2) - (4 * 4))) + (((5 * (3 + (2 * 4))) + (3 + 4)) / 2))) * (8 + (8 - 2))) + (3 * (11 + 20)))) + (((3 * (18 + 5)) * 2) * (2 * (2 * ((((19 * 5) + ((3 + 4) * 2)) * 5) / 5))))))))) * ((2 * (((((17 * 17) * (3 * 3)) + ((((4 * 4) + ((14 + ((17 * 3) - 4)) * 3)) + (10 * (3 * 17))) * 4)) - (((((4 + ((5 * 5) + 2)) * 3) * 3) + ((19 * (3 * ((2 * ((3 * 8) + 5)) / 2))) + (((((3 * 14) + (9 + 2)) * 5) + (11 + (6 * ((3 * 4) / 2)))) + (((3 * (15 - 2)) + ((3 * ((12 + 8) + ((13 * 2) + (3 * 5)))) + ((((4 * 2) + ((17 * 5) + (((((((9 + 18) - 1) * 3) / 3) / 2) * 2) * 4))) * 2) / 2))) + ((((3 * 2) + 7) + 16) * ((2 * 3) + 7)))))) / 2)) * 2)) + ((2 * ((13 * 7) + ((5 * 5) * 6))) + ((13 * 3) * (((2 * 4) + (2 * (3 * 3))) + 5))))) - (((((4 * 2) + (7 + 4)) * ((11 * 5) + (((((((11 * 4) + ((2 * (2 * 8)) * ((4 + 6) + ((2 + 5) + 6)))) + ((((5 * 3) * ((((((((2 * (6 + 1)) * (3 * 4)) - (((4 * (7 + 1)) - (4 + 3)) * 2)) + (2 * ((((((((((((2 * (((((((15 - 5) * ((((((((2 * ((((((2 * (((7 * 4) * ((((((((((5 + 12) * 4) + ((((((4 + (9 + (8 * 2))) * 7) + humn) / 2) - ((12 * (3 * 10)) + ((18 + ((1 + 10) * (3 + 8))) * 3))) * (3 * ((2 * (6 + ((2 * 3) + 1))) - ((7 + 2) - 2))))) * 2) - ((((11 * (3 * 5)) + ((2 * ((3 + (2 * (3 + 10))) * 2)) * 2)) + (((3 + 4) * 5) * (3 * 5))) + ((8 - 1) * (1 + 6)))) / 11) - ((2 * ((18 + (7 + (2 * (3 * (((3 + 17) / 2) - 1))))) + (2 * 14))) + ((2 * 4) * 4))) + 6) / 7) + (((2 * 16) * 7) + ((((9 + (2 * ((2 * 11) - 3))) * 3) + (2 * ((4 * 4) + ((5 + (8 * 3)) * 3)))) * 2)))) - (19 * (7 * 5)))) + ((2 * (((16 * 2) + (1 + (5 * 2))) + ((4 + (2 + ((3 * 2) * (4 * 3)))) + 1))) * 2)) / 3) + ((((2 * 5) + ((((2 * (4 + ((1 + ((7 + 4) * 2)) * 3))) * (4 + (5 * 3))) + (3 * ((11 * (((2 * 11) / 2) * 2)) + (3 * ((3 + 3) + 1))))) / 7)) - ((7 + ((3 * 2) * (3 * 3))) * 2)) * 2)) / 7) - (((5 * 2) * 2) * (5 * (1 + 8))))) + (((17 * 2) + (3 * 5)) * 15)) * 5) - (((4 * 2) + 3) * 14)) / 3) + ((((4 * 2) * 3) * 12) + ((12 + 1) * (15 + 16)))) / 2) - (((1 + 12) + ((13 + 4) * 2)) * 3))) - (((5 + (3 * 14)) * (3 + 8)) + (((2 * (((2 * ((5 + ((5 * 2) + 16)) * 4)) / 2) + (5 * 5))) / 2) * 2))) / 3) + ((15 + 4) * (5 * 5))) * 2) - (14 * (1 + 7)))) + (12 + 1)) / (3 + 4)) - (((3 * 2) + 1) * (11 + 18))) / 8) + (((((2 * (((13 * 3) / 3) * 2)) + (3 * (5 * 3))) + (((2 * (15 + 2)) + ((((5 * 5) - 8) + (2 * 3)) * 4)) - ((((7 + 2) * 3) * 2) - ((2 * 3) + 11)))) + (((((4 * 12) * (2 * (4 * 4))) / 2) / (3 + 3)) + ((10 - 3) * 5))) + (2 * (((((2 * (1 + 5)) + 5) * 2) + 19) * 4)))) * (4 * 12)) + (2 * (13 + (14 * 5)))) * 2) - ((7 * 13) * 2)) / 6) - ((18 + (5 + (4 * (4 * 3)))) * 2)))) / 3) - ((4 + ((5 + (2 + (3 * 2))) + (2 + 5))) + (((2 * 13) * 5) - 17))) / 5) + (17 * ((1 + ((9 + 2) * 2)) - (3 * 2))))) - (8 * (6 * 3))) * 2)) / 12) - ((((((8 * 2) + 2) + ((5 + 18) + (3 * 2))) * (3 * 5)) / 3) + (((4 + 5) * (5 * (2 + 4))) + ((((3 * (14 + (3 * 7))) / ((2 * 5) - 3)) * 3) * 7)))) + ((3 * (((9 * 5) + (((3 * 15) / 3) + 11)) + (((5 * 5) + (3 * ((((((18 + (3 + 13)) / 2) + 2) * 2) - 1) * 3))) - ((2 * (3 + 4)) * (((2 * 3) + 1) + 1))))) - (((((((4 + (3 * 3)) + 9) + ((3 * 7) * 3)) + ((2 * 13) + ((10 * 5) + ((6 + 1) * 3)))) + (16 * 2)) * 2) / 2))) / 2))) + ((4 + (9 * 3)) + (3 * (((13 + 16) * 2) + 13)))) / 5)) * 2)

		if humn%50000000 == 0 {
			log.Println(humn)
		}
	}

	log.Printf("%d", humn) // "3 int8"
}

func init() {
	rootCmd.AddCommand(day21Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day21Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day21Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
