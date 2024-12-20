package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

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

func findSimilarity(a, b []int) map[int]int {
	// gebruik een for loop om te controleren hoe veel waarden in de andere lijst matchen. bij een match, doe i++. eindwaarde van i is het aantal matches
	allMatches := map[int]int{}
	for _, aNum := range a {
		matches := 0
		for _, bNum := range b {
			if aNum == bNum {
				matches++
			}
			score := aNum * matches
			allMatches[aNum] = score
		}
	}

	return allMatches
}

func findSimilarityScore(x map[int]int) int {
	totalScore := 0
	for _, score := range x {
		totalScore += score
	}
	return totalScore
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
			if err != nil {
				fmt.Println("Error: ", err)
			}
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

		summedDistances := sum(distances)
		fmt.Println("Summed distances: ", summedDistances)

	}

	similarity := findSimilarity(leftSlice, rightSlice)
	fmt.Println(similarity)

	finalScore := findSimilarityScore(similarity)
	fmt.Printf("Similarity score: %d", finalScore)
}
