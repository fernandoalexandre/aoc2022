/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var day int

// fetchCmd represents the fetch command
var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetches the current day's input",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("fetch called")

		path := fmt.Sprintf("inputs/day%d", day)

		os.MkdirAll(path, os.ModePerm)

		file_name := fmt.Sprintf("%s/input", path)

		if fileExists(file_name) {
			file_name = fmt.Sprintf("%s_part2", file_name)
		}

		out, err := os.Create(file_name)
		defer out.Close()

		if err != nil {
			log.Fatal(err)
			return
		}

		session_token, envErr := os.LookupEnv("AOC_TOKEN")

		if !envErr {
			log.Fatal("Missing session token in environment variable AOC_TOKEN")
			return
		}

		client := &http.Client{}

		url := fmt.Sprintf("https://adventofcode.com/2022/day/%d/input", day)
		req, err := http.NewRequest("GET", url, nil)
		req.Header.Add("Cookie", session_token)
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
			return
		}
		defer resp.Body.Close()

		n, err := io.Copy(out, resp.Body)
		fmt.Println(fmt.Sprintf("Fetched and written %d bytes", n))
	},
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func init() {
	rootCmd.AddCommand(fetchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fetchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fetchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	fetchCmd.Flags().IntVarP(&day, "day", "d", 1, "Day to fetch the input for.")
}
