package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func ReadLine(fileName string, handler func(string)) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')

		line = strings.TrimSpace(line)

		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		handler(line)
	}
	return nil
}
func ReadLineByPosition(fileName string, handle func(string), position int) {
	f, err := os.Open(fileName)
	if err != nil {
		return
	}
	pos := 0
	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		pos++
		line = strings.TrimSpace(line)
		if err != nil {
			if err == io.EOF {
				return
			}
			return
		}
		if pos > position && pos <= position+200 {
			handle(line)
		}
	}
	return
}
func ReadLineForCountNum(fileName string) int {
	f, err := os.Open(fileName)
	if err != nil {
		return -2
	}
	buf := bufio.NewReader(f)
	size := 0
	for {
		line, errs := buf.ReadString('\n')

		line = strings.TrimSpace(line)

		size = size + 1
		if errs != nil {
			if errs == io.EOF {
				return size
			}
			return -2
		}
	}

	return size
}
func ReadAll(filePth string) ([]byte, error) {
	f, err := os.Open(filePth)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}
func ReadContentByNODEJSLOG(path string) (string, int) {
	contentbyte, err := ReadAll(path)
	if err != nil {
		return "", -1
	}
	result := string(contentbyte)
	array := strings.Split(result, ";")
	if len(array) < 2 {
		return "", -3
	}
	fileSize := string(array[1])
	results, errs := strconv.Atoi(fileSize)
	if errs != nil {
		return "", -2
	}

	return array[0], results
}

func createFile(path string) error {
	file, error := os.Create(path)
	if error != nil {
		fmt.Println(error)
		return error
	}
	defer file.Close()
	return nil
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
func writeContentToSomeFile(path string, content string) {
	ioutil.WriteFile(path, []byte(content), 0644)

}
