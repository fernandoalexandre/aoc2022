/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"math"
	"math/big"

	"sort"

	"github.com/spf13/cobra"
)

// day11Cmd represents the day11 command
var day11Cmd = &cobra.Command{
	Use:   "day11",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day11 called")

		test_monkeys := []Monkey{}
		test_monkeys = append(test_monkeys, Monkey{[]int{79, 98}, Operation{'*', 19}, 23, 2, 3, 0})
		test_monkeys = append(test_monkeys, Monkey{[]int{54, 65, 75, 74}, Operation{'+', 6}, 19, 2, 0, 0})
		test_monkeys = append(test_monkeys, Monkey{[]int{79, 60, 97}, Operation{'^', 2}, 13, 1, 3, 0})
		test_monkeys = append(test_monkeys, Monkey{[]int{74}, Operation{'+', 3}, 17, 0, 1, 0})

		monkeys := []Monkey{}
		monkeys = append(monkeys, Monkey{[]int{91, 58, 52, 69, 95, 54}, Operation{'*', 13}, 7, 1, 5, 0})
		monkeys = append(monkeys, Monkey{[]int{80, 80, 97, 84}, Operation{'^', 2}, 3, 3, 5, 0})
		monkeys = append(monkeys, Monkey{[]int{86, 92, 71}, Operation{'+', 7}, 2, 0, 4, 0})
		monkeys = append(monkeys, Monkey{[]int{96, 90, 99, 76, 79, 85, 98, 61}, Operation{'+', 4}, 11, 7, 6, 0})
		monkeys = append(monkeys, Monkey{[]int{60, 83, 68, 64, 73}, Operation{'*', 19}, 17, 1, 0, 0})
		monkeys = append(monkeys, Monkey{[]int{96, 52, 52, 94, 76, 51, 57}, Operation{'+', 3}, 5, 7, 3, 0})
		monkeys = append(monkeys, Monkey{[]int{75}, Operation{'+', 5}, 13, 4, 2, 0})
		monkeys = append(monkeys, Monkey{[]int{83, 75}, Operation{'+', 1}, 19, 2, 6, 0})

		d11p1(test_monkeys, 96577, false)
		d11p1(monkeys, 9699690, false)

		d11p1(test_monkeys, 96577, true)
		d11p1(monkeys, 9699690, true)
	},
}

type Operation struct {
	operation rune
	val       int
}

func (a Operation) apply(val int, modulus int) (new_val int) {
	bigval := big.NewInt(int64(val))
	bigopval := big.NewInt(int64(a.val))
	bigmodulus := big.NewInt(int64(modulus))
	switch a.operation {
	case '+':
		new_val = int(bigval.Mod(bigval.Add(bigval, bigopval), bigmodulus).Int64())
	case '*':
		new_val = int(bigval.Mod(bigval.Mul(bigval, bigopval), bigmodulus).Int64())
	case '^':
		new_val = int(bigval.Mod(bigval.Exp(bigval, bigopval, nil), bigmodulus).Int64())
	default:
		log.Fatalf("Operation not supported: %s", string(a.operation))
	}
	return new_val
}

type Monkey struct {
	items          []int
	operation      Operation
	test           int
	if_true_throw  int
	if_false_throw int
	count_inspect  int
}

func d11p1(monkeys []Monkey, modulus int, is_part2 bool) {
	num_turn := 0
	if is_part2 {
		log.Println("Part 2")

		num_turn = 10000
	} else {
		log.Println("Part 1")

		num_turn = 20
	}

	for turn := 0; turn < num_turn; turn++ {
		for monkey_idx := 0; monkey_idx < len(monkeys); monkey_idx++ {
			curr_monkey := monkeys[monkey_idx]
			for item_idx := 0; item_idx < len(curr_monkey.items); item_idx++ {
				// Operation
				new_item := curr_monkey.operation.apply(curr_monkey.items[item_idx], modulus)
				// Divide by 3
				if !is_part2 {
					new_item = int(math.Floor(float64(new_item) / 3))
				}
				// Test
				if new_item%curr_monkey.test == 0 {
					// Throw
					monkeys[curr_monkey.if_true_throw].items = append(monkeys[curr_monkey.if_true_throw].items, new_item)
				} else {
					// Throw
					monkeys[curr_monkey.if_false_throw].items = append(monkeys[curr_monkey.if_false_throw].items, new_item)
				}
			}

			monkeys[monkey_idx].count_inspect += len(curr_monkey.items)
			monkeys[monkey_idx].items = []int{}
		}
		//log.Printf("Round: %d | Monkeys = %v", turn, monkeys)
	}

	inspects := []int{}
	for _, monkey := range monkeys {
		inspects = append(inspects, monkey.count_inspect)
	}

	sort.Ints(inspects)

	log.Printf("Monkeys = %v", monkeys)
	log.Printf("Result = %d (%d * %d) ", inspects[len(inspects)-1]*inspects[len(inspects)-2], inspects[len(inspects)-1], inspects[len(inspects)-2])
}

func init() {
	rootCmd.AddCommand(day11Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day11Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day11Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
