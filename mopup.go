package glub

import (
	"strings"
	"regexp"
)

func Mopup(corpus []string) []string{
	re := regexp.MustCompile("[0-9!-<>:@?/(),\"]+")
	stopWords := ReadFiles([]string{"stopwords.txt"})
	for i , data := range corpus{
		corpus[i] = strings.ToLower(data)
		corpus[i] = re.ReplaceAllString(corpus[i], "")
		corpus[i] = RemoveStopWords(corpus[i] , stopWords)
	}
	return corpus
}

func RemoveStopWords(line string, stopWords []string) string{
	regStr := ""
	for i, word := range stopWords{
		if i != 0{
			regStr += `|`
		}
		regStr += `\Q` +  word +  `\E`
	}
	re := regexp.MustCompile(regStr)
	return re.ReplaceAllString(line, " ")
}

