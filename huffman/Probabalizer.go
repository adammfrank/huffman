package huffman

import (
	"bufio"
	"io"
	"log"
	"strings"
)

// CalculateProbabilities returns an array of RuneProbs, one for each rune
func CalculateProbabilities(input string) map[rune]float64 {
	reader := bufio.NewReader(strings.NewReader(input))

	result := make(map[rune]float64)
	length := float64(0)

	for {
		if c, _, err := reader.ReadRune(); err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		} else if _, prs := result[c]; prs != false {
			result[c]++
			length++

		} else {
			result[c] = 1
			length++
		}
	}

	return result
}
