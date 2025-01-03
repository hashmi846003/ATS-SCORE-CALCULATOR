package matcher

import (
	"strings"
)

func CalculateScore(resume, jobDescription string) float64 {
	resumeWords := strings.Fields(resume)
	jobWords := strings.Fields(jobDescription)

	matches := 0
	for _, word := range resumeWords {
		if contains(jobWords, word) {
			matches++
		}
	}

	return (float64(matches) / float64(len(jobWords))) * 100
}

func contains(words []string, word string) bool {
	for _, w := range words {
		if w == word {
			return true
		}
	}
	return false
}
