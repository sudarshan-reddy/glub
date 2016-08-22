package glub

import (
	"strings"
	"regexp"
	"runtime"
	"sync"
	"fmt" //Temporary status logs
)

func Mopup(corpus []string) []string{

	corpLen := len(corpus)

	runtime.GOMAXPROCS(8)
	var wg sync.WaitGroup
	wg.Add(corpLen)	
	
	stopWords := ReadFiles([]string{"stopwords.txt"})
	stopWords = strings.Split(stopWords[0], ",")
	s := make([]string, corpLen)
	mop1 := mop{corpus: s, stopWords : stopWords}
	for i, data := range corpus{
		go mop1.mopupSlv(&wg, i, data)	
	}
	wg.Wait()
	return mop1.corpus
}

type mop struct{
	corpus []string
	stopWords []string
}

func (m *mop) mopupSlv(wg *sync.WaitGroup, i int, data string) {
	fmt.Println("Starting " , i)
	defer wg.Done()
	re := regexp.MustCompile("[0-9!-<>:@?/(),\"]+")
	m.corpus[i] = strings.ToLower(data)
	m.corpus[i] = re.ReplaceAllString(m.corpus[i], "")
	m.corpus[i] = RemoveStopWords(m.corpus[i] , m.stopWords)
	m.corpus[i] = RemoveExtraWhiteSpaces(m.corpus[i])
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

func RemoveExtraWhiteSpaces(line string) string{
	re := regexp.MustCompile(`^[\s\p{Zs}]+|[\s\p{Zs}]+$`) //trim
	line =  re.ReplaceAllString(line, "")
	return re.ReplaceAllString(line, "")
}
