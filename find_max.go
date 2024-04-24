package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

func rootMax() {
	jsonFile, err := os.Open("hard.json")
	if err != nil {
		fmt.Errorf(err.Error())
	}
	defer jsonFile.Close()
	var arr [][]int
	bytes, err := io.ReadAll(jsonFile)
	if readErr := json.Unmarshal(bytes, &arr); readErr != nil {
		log.Fatal("error unmarshal")
	}
	memo := make([][]*int, len(arr))
	for i := range memo {
		memo[i] = make([]*int, len(arr[i])+1)
	}
	log.Println(findMax(memo, arr, 0, 0))
}

func findMax(memo [][]*int, arr [][]int, row int, col int) int {
	if row >= len(arr) {
		return 0
	}
	if col >= len(arr[row]) {
		return 0
	}
	if memo[row][col] != nil {
		return *memo[row][col]
	}
	sum := max(findMax(memo, arr, row+1, col+1), findMax(memo, arr, row+1, col)) + arr[row][col]
	memo[row][col] = &sum
	return sum
}
