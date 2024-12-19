package main

import "fmt"
import "os"
import "bufio"
import "strings"
import "slices"
import "strconv"

func diff(a, b int) int {
	if a < b {
		return b - a 
	}
	return a - b
}

func sum(nums []int) int {
	total := 0

	for _, num := range nums {
		total += num
	}

	return total
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Fout bij openen bestand: ", err)
	}
	defer file.Close()

	leftSlice := []int{}
	rightSlice := []int{}
	
	r := bufio.NewReader(file)
	for {
		line, _, err := r.ReadLine()
		if len(line) > 0 {
			lineString := string(line)
			split := strings.Split(lineString, "   ")
			splitLeft, err := strconv.Atoi(split[0])
			splitRight, err := strconv.Atoi(split[1])
			if err != nil {
				fmt.Println("Error: ", err)
			}
			leftSlice = append(leftSlice, splitLeft)
			rightSlice = append(rightSlice, splitRight)
			slices.Sort(leftSlice)
			slices.Sort(rightSlice)
		}
		if err != nil {
			fmt.Println(err)
			break
		}
	}

	distances := []int{}
	for i, leftValue := range leftSlice {
		rightValue := rightSlice[i]
		diff := diff(leftValue, rightValue)
		distances = append(distances, diff)
	}

	summedDistances := sum(distances)

	fmt.Println("Summed distances: ", summedDistances)
}
