package glub

import (
	"text/scanner"
	"strings"
	"sort"
)

func SortByVal(wordFrequencies map[string]int) PairList{
  pl := make(PairList, len(wordFrequencies))
  i := 0
  for k, v := range wordFrequencies {
    pl[i] = Pair{k, v}
    i++
  }
  sort.Sort(sort.Reverse(pl))
  return pl
}

/*overriding sort*/
type Pair struct {
  Key string
  Value int
}

type PairList []Pair

func (p PairList) Len() int {
	return len(p) 
}

func (p PairList) Less(i, j int) bool {
	return p[i].Value < p[j].Value
}

func (p PairList) Swap(i, j int){
	p[i], p[j] = p[j], p[i]
}

func GenTokens(inputset []string) (map[string]int, []map[string]int) {
	var dataset []map[string]int
	allTokens := make(map[string]int)
	for _, emails := range inputset{
		tokens := make(map[string]int)
		reader := strings.NewReader(emails)
		var scn scanner.Scanner
		scn.Init(reader)
		tok := scn.Scan()
		tokens[scn.TokenText()]++ 
		allTokens[scn.TokenText()] = 0
		for tok != scanner.EOF{
			tok = scn.Scan()
			tokens[scn.TokenText()]++
			allTokens[scn.TokenText()] = 0
		}
		dataset = append(dataset,tokens)
	}
	return allTokens, dataset	
}

func Prep(tokens map[string]int, metadata []map[string]int) [][]int{
	var allWords [][]int
	for word , _ := range tokens{
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
	return allWords
}
