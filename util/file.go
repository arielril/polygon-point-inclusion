package util

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// ParsedFile is the type of the parsed file information
type ParsedFile [][]float32

const coordinatesDivisor = ";"

func createDataObject(data string) ParsedFile {
	lines := strings.Split(data, "\n")

	var dataObject ParsedFile

	for i := 0; i < len(lines); i++ {
		if lines[i] == "" {
			continue
		}

		coordinates := strings.Split(lines[i], coordinatesDivisor)

		var coordLine []float32
		for _, v := range coordinates {
			num, _ := strconv.ParseFloat(v, 32)
			coordLine = append(coordLine, float32(num))
		}

		dataObject = append(dataObject, coordLine)
	}

	return dataObject
}

// ParseFile parses the informed file and creates the data object from the file
func ParseFile(filename string) ParsedFile {
	data, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(fmt.Errorf("failed to open file %v. Error: %v", filename, err))
	}

	return createDataObject(string(data))
}
