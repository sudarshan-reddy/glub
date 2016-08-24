package glub

import (
	"text/scanner"
	"strings"
)

func GenTokens(inputset []string) ([]string, []map[string]int) {
	var dataset []map[string]int
	var allTokens []string
	for _, emails := range inputset{
		tokens := make(map[string]int)
		reader := strings.NewReader(emails)
		var scn scanner.Scanner
		scn.Init(reader)
		tok := scn.Scan()
		tokens[scn.TokenText()]++ 
		allTokens = append(allTokens, scn.TokenText())
		for tok != scanner.EOF{
			tok = scn.Scan()
			tokens[scn.TokenText()]++
			allTokens = append(allTokens, scn.TokenText())
		}
		dataset = append(dataset,tokens)
	}
	return allTokens, dataset	
}

func Prep(tokens []string, metadata []map[string]int, spam int) [][]int{
	var newSet []int
	var allWords [][]int
	for _ , values := range metadata{
		newSet = []int{}
		for _ , words := range tokens{
			newSet = append(newSet, values[words])
		}
		newSet = append(newSet, spam)
		allWords = append(allWords, newSet)
	}
	return allWords
}
