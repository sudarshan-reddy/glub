package glub 

import (
	"fmt"
	"os"
	"path/filepath"
	"io/ioutil"
	"regexp"
	"strings"
)

func GetFiles(dir string) []string {
	fileList := []string{}
	filepath.Walk(dir, func(path string, f os.FileInfo, err error)error{
		if !f.IsDir(){
			fileList = append(fileList, path)
		}
		return nil
	})
	return fileList
}

func ReadFiles(fileList []string) []string {

	corpus := []string{}
	regStr := ""

	for i, word := range returnStopWords(){
		if i != 0{
			regStr += `|`
		}
		regStr += `\b\Q` +  word +  `\E\b`
	}
	re := regexp.MustCompile(regStr + "|[0-9!-<>':@?/(),\"]+")

	for _, files := range fileList{
		body, err := ioutil.ReadFile(files)
		if err != nil{
			fmt.Println(err)
		}
		data := strings.ToLower(string(body))
		data = re.ReplaceAllString(data, "")
		
		corpus = append(corpus, data)
	}
	return corpus
}


