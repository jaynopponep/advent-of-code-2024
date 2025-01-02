package main

import (
	"os"
	"bufio"
)

func main() {
	// idea here is to use DFS when an X is identified, similar to num of islands problem
	// as long as we move left to the right, top bottom, we won't overcount as long as we start
	// DFS searching when we identify an X. 
	f, err := os.Open("xmas_test.txt")
	if err != nil { fmt.Println("could not open xmas txt file", err) 
	return }
	defer f.Close()
	scanner := bufio.NewScanner(f)

}
