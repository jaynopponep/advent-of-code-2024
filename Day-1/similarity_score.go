package main

import (
	"os"
	"bufio"
	"strings"
	"fmt"
	"strconv"
)

func main() {
	f, _ := os.Open("locations.txt")
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	m := make(map[string]int)
	for scanner.Scan() {
		line := scanner.Text()
		right := string(line[len(line)/2:])
		right_t := strings.TrimSpace(right)
		m[right_t] = m[right_t] + 1 
	}
	f.Seek(0,0)
	scanner = bufio.NewScanner(f)
	similarity_score := 0
	for scanner.Scan() {
		line := scanner.Text()
		left := string(line[0:len(line)/2])
		left_t := strings.TrimSpace(left)
		left_int, err := strconv.Atoi(left_t)
		if err != nil { fmt.Println("Error converting number", err)
		return }
		sum := left_int * m[left_t]
		similarity_score += sum
	}
	fmt.Println(similarity_score)
}
