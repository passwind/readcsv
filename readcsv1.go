package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"

	"golang.org/x/text/encoding/simplifiedchinese"
)

func CheckAndEncodeUTF8(filename string) (bool, string, error) {
	result := []string{}
	utf8flag := true

	f, _ := os.Open(filename)

	bfreader := bufio.NewReader(f)
	d := simplifiedchinese.GBK.NewDecoder()
	for {
		line, _, err := bfreader.ReadLine()
		if err == io.EOF {
			break
		}

		if len(line) > 0 {
			utf8f := utf8.Valid(line)
			rowItem := string(line)
			if !utf8f {
				utf8flag = false
				if ret, err := d.String(rowItem); err == nil {
					rowItem = ret
				} else {
					fmt.Printf("err: %s\n", err)
					return false, "", fmt.Errorf("error to decode: %s", err)
				}
			}

			result = append(result, rowItem)

			// TODO: debug
			// record := strings.Split(rowItem, ",")
			// fmt.Println(record)
			// fmt.Println(len(record))
			// for value := range record {
			// 	fmt.Printf("value: %d  %v\n", value, record[value])
			// }
		}
	}

	return utf8flag, strings.Join(result, "\r\n"), nil
}

func main() {
	// Load a TXT file.
	// f, _ := os.Open("/Users/zhuyu/workspace/webapp/test.csv")
	fn := "/Users/zhuyu/workspace/webapp/go/src/g.whispir.cc/service/shopee-whopstool/sample/ordersamplecsv.csv"
	// fn := "/Users/zhuyu/workspace/webapp/go/src/g.whispir.cc/service/shopee-whopstool/sample/ordersamplecsvgb.csv"
	utf8flag, content, err := CheckAndEncodeUTF8(fn)
	fmt.Println(utf8flag, err)
	if err != nil {
		fmt.Printf("err: %s", err)
		return
	}

	var r *csv.Reader

	if !utf8flag {
		r = csv.NewReader(strings.NewReader(content))
	} else {
		f, _ := os.Open(fn)
		r = csv.NewReader(bufio.NewReader(f))
	}

	row := 0

	items := [][]string{}

	for {
		record, err := r.Read()
		// Stop at EOF.
		if err == io.EOF {
			break
		}
		if row == 0 {
			row++
			continue
		}
		// Display record.
		// ... Display record length.
		// ... Display all individual elements of the slice.
		// fmt.Printf("%d %d %#v\n", row, len(record), record)
		fmt.Printf("%d %d\n", row, len(record))
		// for value := range record {
		// 	fmt.Printf("value: %d  %v\n", value, record[value])
		// }
		var subrecord []string
		// parent_key 2
		// items begin from 20, length 10
		var k int
		var emptyFlag bool
		for i, v := range record {
			if i+1 < 20 {
				continue
			} else if (i+1-20)%10 == 0 {
				subrecord = make([]string, 10+1)
				subrecord[0] = record[2-1]
				k = 1
				emptyFlag = true
			}
			if v != "" {
				emptyFlag = false
			}
			subrecord[k] = v
			k++
			if k == 10+1 {
				if emptyFlag {
					break
				}
				items = append(items, subrecord)
			}
		}
		row++
	}

	fmt.Printf("%d %#v\n", len(items), items)
}
