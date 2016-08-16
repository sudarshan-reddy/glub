package main

import (
	"glub"
	"fmt"
)

func main(){
	path := "C:\\Users\\sudarsang\\Downloads\\enron1\\enron1\\enron1\\test"
	fileList := glub.GetFiles(path)
	content := glub.ReadFiles(fileList)
	content = glub.Mopup(content)
	fmt.Println(content)

}
