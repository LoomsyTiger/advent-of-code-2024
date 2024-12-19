package main

import "fmt"
import "os"
import "bufio"
import "strings"
import "slices"

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Fout bij openen bestand: ", err)
	}
	defer file.Close()

	leftSlice := []string{}
	rightSlice := []string{}
	
	r := bufio.NewReader(file)
	for {
		line, _, err := r.ReadLine()
		if len(line) > 0 {
			lineString := string(line)
			split := strings.Split(lineString, "   ")
			leftSlice = append(leftSlice, split[0])
			rightSlice = append(rightSlice, split[1])
			slices.Sort(leftSlice)
			slices.Sort(rightSlice)
			fmt.Println("Left: ", leftSlice)
			fmt.Println("Right: ", rightSlice)
		}
		if err != nil {
			fmt.Println(err)
			break
		}
	}
}
