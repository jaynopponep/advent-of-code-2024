package main

import (
	"fmt"
	"regexp"
	"os"
	"bufio"
	"strconv"
)

func main() {
	// regexp library helps identify basically anything that looks like the following:
	// mul(1,4), anything starting with mul(, followed by any expression [0-9]+, 
	// followed by a comma, then [0-9]+, ending in a ).
	// regexp string that works after testing with .MatchString() is: "mul\\([0-9]+,[0-9]+\\)"
	r, _ := regexp.Compile(`mul\(([0-9]+),([0-9]+)\)`)
	f, err := os.Open("corrupted_muls.txt")
	if err != nil { fmt.Println("could not open the file", err) 
	return }
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var total int
	for scanner.Scan() {
		line := scanner.Text()
		// scan each line, i noticed there were only like 5-6 lines in total, and cutoffs did not interrupt any mul regex
		regex_found := r.FindAllString(line, -1)
		// regex_found is an array of all strings that match the regex, a for loop is needed
		// to use FindStringSubmatch for each array element to get the two multiplicand values
		for i := 0; i < len(regex_found); i++ {
			matched_vals := r.FindStringSubmatch(regex_found[i])
			num1, err := strconv.Atoi(matched_vals[1])
			if err != nil { fmt.Println("couldn't convert match val 1", err)}
			num2, err := strconv.Atoi(matched_vals[2])
			if err != nil { fmt.Println("couldn't convert match val 2", err)}
			product := num1 * num2
			total += product
		}
	}

	if err := scanner.Err(); err != nil { fmt.Println("error during scanning", err)}
	fmt.Println(total)
}
