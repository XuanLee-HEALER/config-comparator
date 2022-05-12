package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

const (
	blockSize = 1024
)

func main() {
	file1 := os.Args[1]
	res := parseFile(file1)
	fmt.Printf("comment: \n%v", res)
}

func parseFile(fileName string) map[interface{}]interface{} {
	fData := readFile(fileName)
	res := make(map[interface{}]interface{})

	err := yaml.Unmarshal(fData, res)
	if err != nil {
		log.Fatalf("解析yaml文件错误，错误信息%v", err)
	}
	return res
}

func readFile(fileName string) []byte {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("打开文件%v错误，错误信息：%v", fileName, err)
	}
	defer f.Close()
	fData := make([]byte, blockSize)
	r, err := f.Read(fData)
	if err != nil {
		log.Fatalf("读取文件%v失败，错误信息：%v", fileName, err)
	}
	if len(fData) <= blockSize {
		fData = fData[:r]
	}

	return fData
}
