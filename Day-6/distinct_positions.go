package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
)

type Tuple struct {
    R int
    C int
}

type TupleSet map[Tuple]struct{}

func main() {
    guards_arr := txt_to_rune("guard.txt")
    guard_location := findGuard(guards_arr)
    guards_arr[guard_location[0]][guard_location[1]] = "."
    distinct_locations := getLocations(guard_location, guards_arr)
    fmt.Println("guard location:", guard_location)
    fmt.Println("distinct locations: ", distinct_locations)
}

func getLocations(guard_location []int, guards_arr [][]string) int {
    up, right, down, left, oob := true, false, false, false, false
    set := make(TupleSet)
    curr := []int{guard_location[0], guard_location[1]}
    distinct_locations := 0
    for !oob {
        for up {
            //fmt.Println("current upwards: ", curr[0], curr[1])
            if curr[0] == 0 {
                distinct_locations += 1
                oob = true
                return distinct_locations
            }
            if ((guards_arr[curr[0]-1][curr[1]]) == "."){
                tuple := Tuple{R: curr[0], C: curr[1]}
                if _, exists := set[tuple]; exists {
                    curr[0] = curr[0]-1
                    continue
                } else {
                    set[Tuple{R: curr[0], C: curr[1]}] = struct{}{}
                    distinct_locations += 1
                }
                curr[0] = curr[0]-1
            } else {
                right = true
                up = false
            }
        }
        for right {
            //fmt.Println("current right: ", curr[0], curr[1])
            if curr[1] == len(guards_arr[0])-1 {
                distinct_locations += 1
                oob = true
                return distinct_locations
            }
            if ((guards_arr[curr[0]][curr[1]+1]) == ".") {
                tuple := Tuple{R: curr[0], C: curr[1]}
                if _, exists := set[tuple]; exists {
                    curr[1] = curr[1]+1
                    continue
                } else {
                    set[Tuple{R: curr[0], C: curr[1]}] = struct{}{}
                    distinct_locations += 1
                }
                curr[1] = curr[1]+1
            } else {
                down = true
                right = false
            }
        }
        for down {
            //fmt.Println("current down: ", curr[0], curr[1])
            if curr[0] == len(guards_arr)-1 {
                distinct_locations += 1
                oob = true
                return distinct_locations
            }
            if ((guards_arr[curr[0]+1][curr[1]]) == ".") {
                tuple := Tuple{R: curr[0], C: curr[1]}
                if _, exists := set[tuple]; exists {
                    curr[0] = curr[0]+1
                    continue
                } else {
                    set[Tuple{R: curr[0], C: curr[1]}] = struct{}{}
                    distinct_locations += 1
                }
                curr[0] = curr[0]+1
            } else {
                left = true
                down = false 
            }
        }
        for left {
            //fmt.Println("current left: ", curr[0], curr[1])
            if curr[1] == 0 {
                distinct_locations += 1
                oob = true
                return distinct_locations
            }
            if ((guards_arr[curr[0]][curr[1]-1]) == ".") {
                tuple := Tuple{R: curr[0], C: curr[1]}
                if _, exists := set[tuple]; exists {
                    curr[1] = curr[1]-1
                    continue
                } else {
                    set[Tuple{R: curr[0], C: curr[1]}] = struct{}{}
                    distinct_locations += 1
                }
                curr[1] = curr[1]-1
            } else {
                up = true
                left = false 
            }
        }
    }
    return 0
}

func txt_to_rune(filename string) [][]string {
    var guards_arr [][]string
    f, err := os.Open(filename)
    if err != nil {
        fmt.Println("error opening file", err)
        return nil 
    }
    defer f.Close()
    scanner := bufio.NewScanner(f)
    scanner.Split(bufio.ScanLines)
    for scanner.Scan() {
        line := scanner.Text()
        fmt.Println("line: ", line)
        words := strings.Split(line, "")
        var line_arr []string
        for index := range words {
            line_arr = append(line_arr, words[index])
        }
        guards_arr = append(guards_arr, line_arr)
    }
    return guards_arr
}

func findGuard(guards_arr [][]string) []int {
    for row:=0; row<len(guards_arr); row++ {
        for col:=0; col<len(guards_arr[row]); col++ {
            if guards_arr[row][col] == "^" {
                guard_location := []int{row, col}
                return guard_location
            }
        }
    }
    return nil
}
