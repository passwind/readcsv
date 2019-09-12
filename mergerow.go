package main

import "fmt"

func MergeRow(orig [][]string, headers []string, indexColStart, indexColEnd, subColStart, subColEnd int) ([][]string, int, int, error) {
	iheaders := []string{}
	dheaders := []string{}
	if len(headers) > 0 {
		for i, hcol := range headers {
			if i+1 >= indexColStart && i+1 <= indexColEnd {
				iheaders = append(iheaders, hcol)
			} else if i+1 >= subColStart && i+1 <= subColEnd {
				dheaders = append(dheaders, hcol)
			}
		}
	}

	tmp := make(map[string][]string)
	for _, orow := range orig {
		rowkey := ""
		indexrow := []string{}
		for j, ocol := range orow {
			if j+1 >= indexColStart && j+1 <= indexColEnd {
				rowkey = rowkey + ocol
				indexrow = append(indexrow, ocol)
			}
		}
		var nrow []string
		if v, ok := tmp[rowkey]; ok {
			nrow = v
		} else {
			nrow = indexrow
		}
		for j, ocol := range orow {
			if j+1 >= subColStart && j+1 <= subColEnd {
				nrow = append(nrow, ocol)
			}
		}
		tmp[rowkey] = nrow
	}
	data := [][]string{}
	var row, col int
	for _, v := range tmp {
		data = append(data, v)
		if len(v) > col {
			col = len(v)
		}
	}
	row = len(data)

	rheaders := []string{}
	for _, v := range iheaders {
		rheaders = append(rheaders, v)
	}
	for i := 0; i < (col-(indexColEnd-indexColStart+1))/(subColEnd-subColStart+1); i++ {
		for _, v := range dheaders {
			rheaders = append(rheaders, fmt.Sprintf("%s_%d", v, (i+1)))
		}
	}

	results := [][]string{}
	results = append(results, rheaders)
	for _, v := range data {
		results = append(results, v)
	}

	return results, row, col, nil
}

func main() {
	records := [][]string{
		[]string{"test1", "test2", "head1", "123"},
		[]string{"test1", "test2", "head2", "123"},
		[]string{"test1", "test2", "head3", "123"},
		[]string{"test1", "test2", "head4", "123"},
		[]string{"test1", "test3", "head1", "123"},
		[]string{"test1", "test3", "head2", "123"},
		[]string{"test2", "test3", "head2", "123"},
	}
	results, row, col, err := MergeRow(records, []string{"id", "name", "item", "amount"}, 1, 2, 3, 4)
	fmt.Println("debug xxxxxx", row, col, err)
	for i, drow := range results {
		fmt.Printf("%d: ", i)
		for _, dcol := range drow {
			fmt.Printf("%s, ", dcol)
		}
		fmt.Printf("\n")
	}
}
