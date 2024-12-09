package Go_Util_AoC

import (
	"bufio"
	"log"
	"os"
)

func ReadInput(inputFile string) []string {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var inputArray []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputArray = append(inputArray, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return inputArray
}

func AbsInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
