package glub

import (
	"text/scanner"
	"strings"
)

func GenTokens(inputset []string , tokens map[string]int) map[string]int {
	for _, emails := range inputset{
		reader := strings.NewReader(emails)
		var scn scanner.Scanner
		scn.Init(reader)
		tok := scn.Scan()
		tokens[scn.TokenText()]++ 
		for tok != scanner.EOF{
				tok = scn.Scan()
				tokens[scn.TokenText()]++
		}
	}
	return tokens	
}
