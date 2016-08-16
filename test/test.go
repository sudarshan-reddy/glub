package main

import (
	"glub"
	"fmt"
)

func main(){
	path := "C:\\Users\\sudarsang\\Downloads\\enron1\\enron1\\enron1"
	fileList := glub.GetFiles(path)
	content := glub.ReadFiles(fileList)
	fmt.Println(content)
}
