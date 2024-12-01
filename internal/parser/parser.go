package parser

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func ParseTxt(filePath string) ([]string, error) {
	data, err := getFileData(filePath)
	if err != nil {
		return nil, errors.New("failed to read file: " + filePath + ". " + err.Error())
	}

	lines := strings.Split(string(data), "\n")
	return lines, nil
}

func getFileData(fileName string) ([]byte, error) {
	fmt.Println("Reading " + fileName)
	data, err := os.ReadFile(fileName)

	return data, err
}
