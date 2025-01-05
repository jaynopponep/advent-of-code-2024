package main

import (
	"os"
	"bufio"
	"fmt"
)

var xmas_arr [][]rune

func verify_xmas(origin []int, direction []int) bool {
	// expects origin and directions like [-1, -1], [0, 1], etc
	// returns true or false
	// example: verify_xmas([]int{4,0}, []int{1, 1}) 
	verified, M, A := false, false, false
	curr := []int{origin[1], origin[0]}
	var delta_x, delta_y int
	for !verified {
		delta_x = curr[0]+direction[0]
		delta_y = curr[1]+direction[1]
		if 0 > delta_x || len(xmas_arr[0]) <= delta_x {
			verified = false
			break
		}
		if 0 > delta_y || len(xmas_arr) <= delta_y {
			verified = false
			break
		}
		curr = []int{delta_x, delta_y}
		letter := xmas_arr[curr[1]][curr[0]]
		if A && letter == 'S' {
			verified = true
			//fmt.Println("Verified following xmas coords:", origin[0], origin[1])
			break
		} else if M && letter == 'A' && A == false{
			A = true
			continue
		} else if letter == 'M' && M == false{
			M = true
			continue
		} else {
			verified = false
			break
		}
	}
	return verified
}

func main() {
	// tasks:
	// - read each line to become an array of strings
	// - split each element into array of characters, creating a 2d array as a result
	// - now you have a 2d map
	// - write subroutines that scan if an X becomes a full XMAS based on direction given
	// specification:
	// verify_xmas(direction) -> true or false
	// direction can be: [-1,-1], [0,-1], [0,1], etc. we only need one direction because
	// all following letters will follow the same direction given.
	// indicator for calling verify_xmas is when one of the directions is true for when checking if
	// the -1,-1 of an X is an M. no need to worry for double counts because we will only check when
	// we see an X, and we won't re-check X's if we go left right, top bottom.
	f, err := os.Open("xmas.txt")
	if err != nil { fmt.Println("could not open xmas txt file", err) 
	return }
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		line_arr := []rune(line)
		xmas_arr = append(xmas_arr, line_arr)
	}

	// show the chars; it is correct chars in the []rune, but will not display as such
	// unless one does string() on all rows:
	count := 0
	for y := 0; y < len(xmas_arr); y++ {
		for x := 0; x < len(xmas_arr[0]); x++ {
			if xmas_arr[y][x] == 'X' {
				if verify_xmas([]int{y, x}, []int{-1, 0}) {
					count += 1
				}
				if verify_xmas([]int{y, x}, []int{1, 0}) {
					count += 1
				}
				if verify_xmas([]int{y, x}, []int{0, 1}) {
					count += 1
				}
				if verify_xmas([]int{y, x}, []int{0, -1}) {
					count += 1
				}
				if verify_xmas([]int{y, x}, []int{1, 1}) {
					count += 1
				}
				if verify_xmas([]int{y, x}, []int{1, -1}) {
					count += 1
				}
				if verify_xmas([]int{y, x}, []int{-1, 1}) {
					count += 1
				}
				if verify_xmas([]int{y, x}, []int{-1, -1}) {
					count += 1
				}
			}
		}
	}
	fmt.Println("count total: ", count)
}
// I'm fully aware the above code is pretty unreadable. it works though! : D
