/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"log"

	"github.com/fernandoalexandre/aoc2022/cmd"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cmd.Execute()
}
