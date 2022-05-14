package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

const (
	blockSize = 8
)

type Entries map[interface{}]interface{}

func main() {
	file1 := os.Args[1]
	var entries Entries
	err := parseFile(file1, &entries)
	if err != nil {
		log.Printf("比较失败，错误信息：%v", err)
	}

	handleEntries(&entries)
}

func handleEntries(entries *Entries) {
	for k, _ := range *entries {
		switch k.(type) {
		case string:
			fmt.Println(k)
		default:
			break
		}
	}
}

func parseFile(fileName string, entries *Entries) (err error) {
	fData, err := readFile(fileName)
	if err != nil {
		log.Printf("打开文件%v错误，错误信息：%v", fileName, err)
		return err
	}

	err = yaml.Unmarshal(fData, entries)
	if err != nil {
		log.Printf("解析yaml文件错误，错误信息：%v", err)
		return err
	}
	return nil
}

func readFile(fileName string) (res []byte, err error) {
	f, err := os.Open(fileName)
	if err != nil {
		log.Printf("打开文件%v错误，错误信息：%v", fileName, err)
	}
	defer f.Close()
	buf := make([]byte, blockSize)
	res, rc := make([]byte, 0, blockSize), 0
	for {
		r, err := f.Read(buf)
		res = append(res, buf[:r]...)
		rc += r
		if err != nil && err != io.EOF {
			log.Printf("读取文件%v失败，错误信息：%v", fileName, err)
			return nil, err
		}

		if err == io.EOF {
			break
		}
	}

	res = res[:rc]

	return res, nil
}
