package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	// Load a TXT file.
	// f, _ := os.Open("/Users/zhuyu/workspace/webapp/test.csv")
	f, _ := os.Open("/Users/zhuyu/workspace/webapp/go/src/g.whispir.cc/service/shopee-whopstool/sample/ordersamplecsv.csv")

	// Create a new reader.
	r := csv.NewReader(bufio.NewReader(f))
	for {
		record, err := r.Read()
		// Stop at EOF.
		if err == io.EOF {
			break
		}
		// Display record.
		// ... Display record length.
		// ... Display all individual elements of the slice.
		fmt.Println(record)
		fmt.Println(len(record))
		for value := range record {
			fmt.Printf("value: %d  %v\n", value, record[value])
		}
	}
}
