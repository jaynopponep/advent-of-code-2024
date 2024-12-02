package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"math"
	"strconv"
)

func main() {
	f, _ := os.Open("reports.txt")
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	safe := 0
	for scanner.Scan() {
		safety := true 
		line := scanner.Text()
		words := strings.Fields(line)
		var upwards_dir bool
		for index := range words {
			if index == 0 {
				second, err := strconv.Atoi(words[1])
				if err != nil { fmt.Println("Trouble converting curr num", err) 
				return}
				first, err := strconv.Atoi(words[0])
				if err != nil { fmt.Println("Trouble converting curr num", err) 
				return}
				diff := second - first
				if diff > 0 {
					upwards_dir = true
				} else {
					upwards_dir = false 
				}
			} else {
				second, err := strconv.Atoi(words[index])
				if err != nil { fmt.Println("Trouble converting curr num", err) 
				return}
				first, err := strconv.Atoi(words[index-1])
				if err != nil { fmt.Println("Trouble converting curr num", err) 
				return}
				diff := second - first
				diff_abs := math.Abs(float64(diff))
				if diff == 0 {
					safety = false
					break
				} else if diff_abs > 3 {
					safety = false
					break
				} else if diff > 0 {
					if upwards_dir == false {
						safety = false
						break
					}
				} else {
					if upwards_dir == true {
						safety = false
						break
					}
				}
			}
		}
		if safety == true {
			safe += 1
		}
	}
	fmt.Println(safe)

}
