package glub 

import (
	"fmt"
	"os"
	"path/filepath"
	"io/ioutil"
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
	for _, files := range fileList{
		body, err := ioutil.ReadFile(files)
		if err != nil{
			fmt.Println(err)
		}
		corpus = append(corpus, string(body))
	}
	return corpus
}


