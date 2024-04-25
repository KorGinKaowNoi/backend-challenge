package main

import (
	"fmt"
)

func leftRight() {
	var encode string
	_, err := fmt.Scanln(&encode)
	if err != nil {
		return
	}
	result := make([]int, len(encode)+1)
	sEncode := []rune(encode)
	for i, char := range encode {
		if char == 'L' {
			for j := i; j >= 0; j-- {
				if sEncode[j] == 'L' {
					getL(result, j, j+1)
				} else if sEncode[j] == 'R' {
					getR(result, j, j+1)
				} else {
					getE(result, j, j+1)
				}
			}
		} else if char == 'R' {
			getR(result, i, i+1)
		} else {
			getE(result, i, i+1)
		}
		fmt.Println(result)
	}
	fmt.Println(result)
}

func getL(arr []int, left int, right int) {
	if arr[left] <= arr[right] {
		arr[left] = arr[right] + 1
	}
}

func getR(arr []int, left int, right int) {
	if arr[right] <= arr[left] {
		arr[right] = arr[left] + 1
	}
}

func getE(arr []int, left int, right int) {
	m := max(arr[left], arr[right])
	arr[left] = m
	arr[right] = m
}
