package glub

import (
	"strings"
	"regexp"
)

func Mopup(corpus []string) []string{
	re := regexp.MustCompile("[0-9!-<>:@?/(),\"]+")
	for i , data := range corpus{
		corpus[i] = strings.ToLower(data)
		corpus[i] = re.ReplaceAllString(corpus[i], "")
	}
	return corpus
}

