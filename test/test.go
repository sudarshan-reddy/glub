package main

import (
	"glub"
	"fmt"
)

func main(){
	fmt.Println("test")
	path := "C:\\Users\\sudarsang\\Downloads\\enron1\\enron1\\enron1"
	fileList := glub.GetFiles(path)
	content := glub.ReadFiles(fileList)
	content = glub.Mopup(content)
	tokens := make(map[string]int)
	tokens = glub.GenTokens(content, tokens)
	for k, v := range tokens{
		fmt.Println(k, v)
	}

}
