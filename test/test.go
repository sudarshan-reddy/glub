package main

import (
	"glub"
	"fmt"
	"time"
)

func main(){
	fmt.Println("test")
	start := time.Now()
	path := "C:\\Users\\sudarsang\\Downloads\\enron1\\enron1\\enron1"
	fileList := glub.GetFiles(path)
	content := glub.ReadFiles(fileList)
	fmt.Println("mopping up")
	content = glub.Mopup(content)
	tokens := make(map[string]int)
	tokens = glub.GenTokens(content, tokens)
	for k, v := range tokens{
		fmt.Println(k, v)
	}
	elapsed := time.Since(start)
	fmt.Println("time taken: ",  elapsed)

}
