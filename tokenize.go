package glub

import (
	"text/scanner"
	"strings"
	"strconv"
	"fmt"
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

func Prep(tokens []string, metadata []map[string]int, spam int) [][]string{
	var newSet []string
	var allWords [][]string
	for _ , values := range metadata{
		newSet = []string{}
		for _ , words := range tokens{
			newSet = append(newSet, strconv.Itoa(values[words]))
		}
		newSet = append(newSet, strconv.Itoa(spam))
		allWords = append(allWords, newSet)
	}
	return allWords
}
