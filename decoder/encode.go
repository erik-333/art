package main

import (
	"fmt"
	"strings"
)

// this function works as following, it checks length of repeating sequences by checking how many times ;
// you can repeat a certain length, it keeps checking if longer sequences are possible, if not it will ;
// output the current sequence, if it finds a longer sequence it will output that instead. :DD ;
func encode(input string) string {
	var result strings.Builder
	runes := []rune(input)
	n := len(runes)

	for i := 0; i < n; {
		// Max possible sequence length
		maxSeqLen := (n - i) / 2
		// Length of the best repeating sequence found
		bestSeqLen := 1
		// Number of repetitions for the best sequence
		bestRepCount := 1

		// Try sequences of different lengths
		for seqLen := 1; seqLen <= maxSeqLen; seqLen++ {
			seq := string(runes[i : i+seqLen])
			count := 1
			pos := i + seqLen

			// Count how many times this sequence repeats
			for pos+seqLen <= n && string(runes[pos:pos+seqLen]) == seq {
				count++
				pos += seqLen
			}

			// Update best sequence if current one is better
			if count > 1 && count*seqLen > bestRepCount*bestSeqLen {
				bestSeqLen = seqLen
				bestRepCount = count
			}
		}

		// If we found a repeating sequence worth compressing
		if bestRepCount > 1 && bestRepCount*bestSeqLen > 3 {
			seq := string(runes[i : i+bestSeqLen])
			result.WriteString(fmt.Sprintf("[%d %s]", bestRepCount, seq))
			i += bestRepCount * bestSeqLen
		} else {
			// No worthwhile repeating sequence found, output single character
			result.WriteRune(runes[i])
			i++
		}
	}

	return result.String()
}
