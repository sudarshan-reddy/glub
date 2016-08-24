package main

import (
	"glub"
	"fmt"
	"time"
)

func main(){
	fmt.Println("test")
	start := time.Now()
	path := "C:\\Users\\sudarsang\\Downloads\\enron1\\enron1\\enron1\\test"
	fileList := glub.GetFiles(path)
	content := glub.ReadFiles(fileList)
	allTokens, metadata := glub.GenTokens(content)
	dataset := glub.Prep(allTokens, metadata)
	for _, vals := range dataset{
		fmt.Println(vals)
	}

	fmt.Println(allTokens)
	elapsed := time.Since(start)
	if dataset != nil{
		fmt.Println("time taken: ",  elapsed)
	}

}
