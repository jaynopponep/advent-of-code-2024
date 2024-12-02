package main

import (
	"fmt"
	"strings"
	"os"
	"bufio"
	"sort"
	"math"
	"strconv"
)

func main() {
	// note split implementation makes sure to split each line in equal halves
	// which means that input must have equal size pair of numbers in each line
	// read file
	f, _ := os.Open("locations.txt")
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	// get # of lines with bufio
	var num_lines int
	for scanner.Scan() {
		num_lines++
	}
	f.Seek(0, 0) // must reset back to beginnig point, after scanner.Scan() pointer goes to the end.
	scanner = bufio.NewScanner(f)
	// create left & right arrs
	var left_nums []string
	var right_nums []string

	for scanner.Scan() {
		// extract left & right nums, append each to their arrays
		line := scanner.Text()
		left := string(line[0:len(line)/2])
		left_t := strings.TrimSpace(left)
		left_nums = append(left_nums, left_t)
		right := string(line[len(line)/2:])
		right_t := strings.TrimSpace(right)
		right_nums = append(right_nums, right_t)
	}
	sort.Strings(left_nums)
	sort.Strings(right_nums)

	// now, we compute after the sorting funcs
	distance := 0
	for i:=0; i<num_lines; i++ {
		left, err := strconv.Atoi(left_nums[i])
		if err != nil { fmt.Println("Error converting left number:", err) 
		return }
		right, err := strconv.Atoi(right_nums[i])
		if err != nil { fmt.Println("Error converting right number:", err) 
		return }
		diff := math.Abs(float64(left - right))
		distance += int(diff)
	}
	fmt.Println(distance)
}
