package hw03_frequency_analysis //nolint:golint,stylecheck
import (
	"sort"
)

const Max = 10

func Top10(input string) []string { // Place your code here
	if len(input) == 0 {
		return make([]string, 0)
	}
	var start = 0
	top := make(map[string]int)
	words := make([]string, 0)

	for pos, char := range input {
		if char == 32 || char == 9 || char == 10 { // if new space or a new line then
			if pos == start {
				start++
				continue
			}
			word := input[start:pos]
			value, ok := top[word]
			top[word] = value + 1
			if !ok {
				words = append(words, word)
			}
			start = pos + 1
		}
	}
	if start < len(input) {
		word := input[start:]
		value, ok := top[word]
		top[word] = value + 1
		if !ok {
			words = append(words, word)
		}
	}

	sort.Slice(words, func(i, j int) bool {
		return top[words[i]] > top[words[j]]
	})
	min := len(words)
	if min > Max {
		min = Max
	}
	output := words[0:min]
	return output
}
