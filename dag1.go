package main

import "fmt"
import "os"
import "bufio"
import "strings"

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
			fmt.Println("Left: ", split[0])
			fmt.Println("Right: ", split[1])
		}
		if err != nil {
			fmt.Println(err)
			break
		}
	}
}
