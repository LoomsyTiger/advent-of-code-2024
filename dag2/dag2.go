package dag2

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Fout bij openen bestand: ", err)
	}
	defer file.Close()

	r := bufio.NewReader(file)
	for {
		line, _, err := r.ReadLine()

		split := []int{}
		split = strings.Split(string(line), " ")

		if err != nil {
			fmt.Printf("Error: %v", err)
		}
	}
}
