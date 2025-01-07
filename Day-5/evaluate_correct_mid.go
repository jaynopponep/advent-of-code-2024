package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("rules.txt")
	if err != nil { fmt.Println("couldn't open file", err)
    return }
	defer f.Close()

	rules_map := make(map[int][]int)
	var print_pages [][]int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" { continue }
		if strings.Contains(line, "|") {
			parts := strings.Split(line, "|")
			key, err1 := strconv.Atoi(strings.TrimSpace(parts[0]))
			value, err2 := strconv.Atoi(strings.TrimSpace(parts[1]))
			if err1 == nil && err2 == nil {
				rules_map[key] = append(rules_map[key], value)
			}
		} else {
			pages := strings.Split(line, ",")
			var pageList []int
			for _, page := range pages {
				pageNum, err := strconv.Atoi(strings.TrimSpace(page))
				if err == nil {
					pageList = append(pageList, pageNum)
				}
			}
			print_pages = append(print_pages, pageList)
		}
	}
	// now to evaluate...
	// 1. if 'curr' value is not a key in the map, it is instantly incorrectly-ordered
	// 2. 'curr' values starts at 0 and end at len(print_pages)-1
	// 3. if curr+1 is in rules_map[curr], then curr & curr+1 are correctly ordered
	var middle_total int
	for _, pages := range print_pages {
		for index, curr_page := range pages[:len(pages)-1] {
			if _, exists := rules_map[curr_page]; !exists {
				break
			} 
			if contains(rules_map[curr_page], pages[index+1]) {
				if index == len(pages)-2 {
					middle_total += pages[len(pages)/2]
				} else {
					continue
				}
			} else {
				break
			}
		}
	}
}

func contains(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
