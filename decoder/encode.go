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
	const maxAllowedSeqLen = 73 // Limit of terminal width, any sequence longer will result in non-optimal compression

	for i := 0; i < n; {
		// Max possible sequence length - take minimum of terminal width and half remaining input
		maxSeqLen := (n - i) / 2
		if maxSeqLen > maxAllowedSeqLen {
			maxSeqLen = maxAllowedSeqLen
		}

		bestSeqLen := 1
		bestRepCount := 1
		bestRatio := 0.0

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

			// Calculate compression ratio
			originalLen := count * seqLen
			compressedLen := len(fmt.Sprintf("[%d %s]", count, seq))
			ratio := float64(originalLen) / float64(compressedLen)

			// Update best sequence if current one has better compression ratio
			if count > 1 && ratio > bestRatio {
				bestSeqLen = seqLen
				bestRepCount = count
				bestRatio = ratio
			}
		}

		// If we found a repeating sequence worth compressing
		if bestRepCount > 1 && bestRatio > 1.25 {
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
