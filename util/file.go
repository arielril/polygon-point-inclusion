package util

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const coordinatesDivisor = ";"

func createDataObject(data string) [][]int {
	lines := strings.Split(data, "\n")

	var dataObject [][]int

	for i := 0; i < len(lines); i++ {
		if lines[i] == "" {
			continue
		}

		coordinates := strings.Split(lines[i], coordinatesDivisor)

		var coordLine []int
		for _, v := range coordinates {
			num, _ := strconv.Atoi(v)
			coordLine = append(coordLine, num)
		}

		dataObject = append(dataObject, coordLine)
	}

	return dataObject
}

// ParseFile parses the informed file and creates the data object from the file
func ParseFile(filename string) interface{} {
	data, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(fmt.Errorf("failed to open file %v. Error: %v", filename, err))
	}

	return createDataObject(string(data))
}
