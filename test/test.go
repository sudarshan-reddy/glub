package main

import (
	"glub"
	"fmt"
	"time"
)

func getOP(filepath string){
	path := "/home/ubuntu/workspace/enrondata/" + filepath
	fmt.Println("getFiles " + filepath)
	fileList := glub.GetFiles(path)
	fmt.Println("readfiles " + filepath)
	content := glub.ReadFiles(fileList)
	fmt.Println("Tokenize" + filepath)
	allTokens, metadata := glub.GenTokens(content)
	fmt.Println("Prep " + filepath)
	dataset := glub.Prep(allTokens, metadata, 0)
	fmt.Println("writeTocsv " + filepath)
	glub.WriteToCsv(allTokens, dataset, filepath + ".csv")
}

func main(){
	fmt.Println("test")
	start := time.Now()
	getOP("ham")
	getOP("spam")
	elapsed := time.Since(start)
	fmt.Println("time: " , elapsed)
}