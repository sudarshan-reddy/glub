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
	allTokens, metadata := glub.GenTokens(content)
	dataset := glub.Prep(allTokens, metadata)
	for keys, vals := range dataset{
		fmt.Println(keys, vals)
	}
	elapsed := time.Since(start)
	fmt.Println("time taken: ",  elapsed)

}
