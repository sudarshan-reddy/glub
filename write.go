package glub

import (
		"os"
		"encoding/csv"
		"fmt"
		)

func WriteToCsv(headers []string, data [][]string , filename string){
	file, err := os.Create(filename)
	errHandle(err)
	defer file.Close()
	writer := csv.NewWriter(file)

	writer.Write(headers)
	for _, values := range data{
		err := writer.Write(values)
		errHandle(err)
	}

	defer writer.Flush()
}

func errHandle(err error) {
    if err != nil {
		fmt.Println(err)
    }
}
