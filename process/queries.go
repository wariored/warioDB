package process

import (
	"fmt"
	"strings"
)

func SelectAll(DbName, tableName string) [][]string {
	tableFilePath := DbPath + DbName + "." + tableName + ".txt"
	lines, err := ScanLines(tableFilePath)
	if err != nil {
		panic(err)
	}
	var lineValuesByArray [][]string
	for _, line := range lines {
		words := strings.Fields(line)
		lineValuesByArray = append(lineValuesByArray, words)
		fmt.Println("---")
		//fmt.Println(line)
	}
	fmt.Println(lineValuesByArray)
	return lineValuesByArray

}
