package main

import (
	"os"
	"bufio"
	"fmt"
)

var xmas_arr [][]rune

func main() {
	f, err := os.Open("x-mas.txt")
	if err != nil { fmt.Println("couldn't open x-mas file", err)
    return }
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		line_arr := []rune(line)
		xmas_arr = append(xmas_arr, line_arr)
	}
	count := 0
	for y := 1; y < len(xmas_arr)-1; y++ { // y -> rows
		for x := 1; x < len(xmas_arr[0])-1; x++ { // x -> columns
			if xmas_arr[y][x] == 'A' {
				//fmt.Println("Found A at row", y, "column", x)
				if verify_x_mas([]int{y, x}) {
					count += 1
				}
			}
		}
	}
	fmt.Println("count total: ", count)
}

func verify_x_mas(origin []int) bool {
	verified, tleft_bright, tright_bleft := false, false, false
	//curr := []int{origin[0], origin[1]}
	//var delta_row, delta_col int
	if origin[0]-1 < 0 || origin[0]+1 >= len(xmas_arr) || origin[1]-1 < 0 || origin[1]+1 >= len(xmas_arr[0]) {
		return false
	}
	for !verified {
		if !tleft_bright {
			if xmas_arr[origin[0]-1][origin[1]-1] == 'M' && xmas_arr[origin[0]+1][origin[1]+1] == 'S' {
				tleft_bright = true
			} else if xmas_arr[origin[0]-1][origin[1]-1] == 'S' && xmas_arr[origin[0]+1][origin[1]+1] == 'M' {
				tleft_bright = true
			} else {
				break
			}
		}
		if !tright_bleft {
			if xmas_arr[origin[0]-1][origin[1]+1] == 'M' && xmas_arr[origin[0]+1][origin[1]-1] == 'S' {
				tright_bleft = true
			} else if xmas_arr[origin[0]-1][origin[1]+1] == 'S' && xmas_arr[origin[0]+1][origin[1]-1] == 'M' {
				tright_bleft = true
			} else {
				break
			}
		}
		if tleft_bright && tright_bleft {
			verified = true
		}
		if !tleft_bright && !tright_bleft { // Add an explicit termination condition
			break
		}
	}
	return verified
}
