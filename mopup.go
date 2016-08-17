package glub

import (
	"strings"
	"regexp"
)

func Mopup(corpus []string) []string{
	re := regexp.MustCompile("[0-9!-<>:@?/(),\"]+")
	stopWords := ReadFiles([]string{"stopwords.txt"})
	stopWords = strings.Split(stopWords[0], ",")
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
		regStr += `\b\Q` +  word +  `\E\b`
	}
	re := regexp.MustCompile(regStr)
	return re.ReplaceAllString(line, "")
}

