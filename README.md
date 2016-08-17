# Simple Corpus creator

Currently only supports text based files from a directory.

Loops through a directory, prunes stop words, certain punctuations and 
tokenizes output into a map.

Installation

```sh
go get github.com/sudsred/glub
```

Usage Example:

```go
package main

import (
	"glub"
	"fmt"
)

func main(){
	path := "full filepath here"
	fileList := glub.GetFiles(path) //gets files
	content := glub.ReadFiles(fileList) //reads files
	content = glub.Mopup(content) //prunes files
	tokens := make(map[string]int) //initializing tokens
	tokens = glub.GenTokens(content, tokens) 
	/*takes tokens and returns them, this is intentional
	to add to existing tokens
	*/
	for k, v := range tokens{
		fmt.Println(k, v)
	}

}
```
