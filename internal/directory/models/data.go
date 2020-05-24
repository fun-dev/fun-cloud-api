package models

import (
	"bufio"
	"strings"
)

type DataType int

const (
	FILE DataType = iota
	DIRECTORY
)

var (
	_suffixDirectoryWord = "/"
)

type Data struct {
	DataType   DataType
	DataPath   string
	DataName   string
	ChildItems []*Data
}

func ParseLsCmdData(lsCmdData, dirPath string) (*Data, error) {
	result := &Data{
		DataType: DIRECTORY,
		DataPath: dirPath,
	}
	oneLiners, err := _parseStringToOneLinerArray(lsCmdData)
	if err != nil {
		return nil, err
	}
	for _, oneLine := range oneLiners {
		data := &Data{
			DataType: _judgeFileOrDirectory(oneLine),
			DataPath: dirPath + "/" + oneLine,
			DataName: strings.Replace(oneLine, _suffixDirectoryWord, "", -1),
		}
		result.ChildItems = append(result.ChildItems, data)
	}
	return result, nil
}

func _judgeFileOrDirectory(data string) DataType {
	isDirectory := strings.HasSuffix(data, _suffixDirectoryWord)
	if isDirectory {
		return DIRECTORY
	}
	return FILE
}

func _parseStringToOneLinerArray(data string) ([]string, error) {
	scanner := bufio.NewScanner(strings.NewReader(data))
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}
