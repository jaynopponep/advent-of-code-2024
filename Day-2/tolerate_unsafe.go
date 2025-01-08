package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"strconv"
	"math"
)

var reports [][]int

func main() {
	// in this file, we use the same as part 1, except we have a bail-out card 
	// variable that can be used once in each iteration only to bypass the 
	// break condition.
	f, _ := os.Open("reports.txt")
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	safe := 0
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		var line_arr []int
		for index := range words {
			num, err := strconv.Atoi(words[index])
			if err != nil { fmt.Println("Trouble converting num", err)
		    return }
			line_arr = append(line_arr, num)
		}
		reports = append(reports, line_arr)
	}
	for index := range reports {
		safety, _ := verify_safe(reports[index])
		if !safety {
			for i:=0; i<len(reports[index]); i++ {
				report_new := make([]int, len(reports[index]))
				copy(report_new, reports[index])
				report_new = append(report_new[:i], report_new[i+1:]...)
				if safety_t, _ := verify_safe(report_new); safety_t {
					safety = true
					break
				}
			}
		}
		if safety {
			safe += 1
		}
	} 
	fmt.Println(safe)
}

func verify_safe(slice []int) (bool, int) {
	safety := true
	var upwards bool
	if slice[1] > slice[0] {
		upwards = true
	} else if slice[1] < slice[0] {
		upwards = false
	} else {
		return false, 0
	}
	for i := 0; i < len(slice)-1; i++ {
		diff_abs := math.Abs(float64(slice[i+1] - slice[i]))
		if diff_abs > 3 || diff_abs < 1 {
			return false, i
		}
		if upwards {
			if slice[i+1] > slice[i] {
				continue
			} else {
				return false, i
			}
		} else {
			if slice[i+1] < slice[i] {
				continue
			} else {
				return false, i
			}
		}
	}
	return safety, len(slice)-1
}
