package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

const (
	blockSize = 8
)

func main() {
	file1 := os.Args[1]

	res, err := parseFile(file1)
	if err != nil {
		log.Printf("比较失败，错误信息：%v", err)
	}

	handleMap(res, 0)
}

func handleMap(parse interface{}, layer int) {
	if f, ok := parse.(map[interface{}]interface{}); ok {
		for k, v := range f {
			if tk, ok := k.(string); ok {
				fmt.Print(strings.Repeat(" ", layer*2), tk, ":")
				switch v.(type) {
				case string, int, bool:
					fmt.Println(" ", v)
				case []interface{}:
					fmt.Print("\n")
					handleArr(v, layer+1)
				case map[interface{}]interface{}:
					fmt.Print("\n")
					handleMap(v, layer+1)
				default:
					fmt.Printf("type of v %T \n", v)
				}
			}
		}
	}
}

func handleArr(arr interface{}, layer int) {
	if f, ok := arr.([]interface{}); ok {
		for _, e := range f {
			handleMap(e, layer)
			fmt.Println()
		}
	}
}

func parseFile(fileName string) (res map[interface{}]interface{}, err error) {
	fData, err := readFile(fileName)
	if err != nil {
		log.Printf("打开文件%v错误，错误信息：%v", fileName, err)
		return nil, err
	}

	res = make(map[interface{}]interface{})
	err = yaml.Unmarshal(fData, res)
	if err != nil {
		log.Printf("解析yaml文件错误，错误信息：%v", err)
		return nil, err
	}
	return res, nil
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
