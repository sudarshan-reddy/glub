package glub

import (
	"text/scanner"
	"strings"
)

func Zip(rawData [][]int) [][]int {
	var transpose [][]int
	for i := 0; i < len(rawData[0]); i++{
		cog := make([]int, len(rawData[0]))
		for _, vals := range rawData {
			cog = append(cog,vals[i])
		}
		transpose = append(transpose, cog)
	}
	return transpose

}

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

func Prep(tokens []string, metadata []map[string]int) [][]int{
	var allWords [][]int
	for _ , word := range tokens{
		var wordSet []int 
		for _ , stuff := range metadata{
			if stuff[word] != 0{
				wordSet = append(wordSet, stuff[word])	
			}else{
				wordSet = append(wordSet, 0)
			}
		}
		allWords = append(allWords, wordSet)
	}
	Zip(allWords)
	return Zip(allWords)
}
