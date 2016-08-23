package glub

import (
	"text/scanner"
	"strings"
)

func GenTokens(inputset []string) []map[string]int {
	dataset := make([]map[string]int, len(inputset))
	for _, emails := range inputset{
		tokens := make(map[string]int)
		reader := strings.NewReader(emails)
		var scn scanner.Scanner
		scn.Init(reader)
		tok := scn.Scan()
		tokens[scn.TokenText()]++ 
		for tok != scanner.EOF{
				tok = scn.Scan()
				tokens[scn.TokenText()]++
		}
		dataset = append(dataset,tokens)
	}
	return dataset	
}
