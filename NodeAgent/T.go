package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

type ProcessStatus struct {
	Name string
	State string
}

func exec_shell(s string) (string, error) {
	//函数返回一个*Cmd，用于使用给出的参数执行name指定的程序
	cmd := exec.Command("/bin/bash", "-c", s)

	//读取io.Writer类型的cmd.Stdout，再通过bytes.Buffer(缓冲byte类型的缓冲器)将byte类型转化为string类型(out.String():这是bytes类型提供的接口)
	var out bytes.Buffer
	cmd.Stdout = &out

	//Run执行c包含的命令，并阻塞直到完成。  这里stdout被取出，cmd.Wait()无法正确获取stdin,stdout,stderr，则阻塞在那了
	err := cmd.Run()
	//checkErr(err)

	return out.String(), err
}

func findPid(name string) {
	_, err := os.Stat("/proc")
	if err == nil {

	}
}

func main() {
	str,err:=exec_shell("ps -ef | grep  node  | grep -v grep | awk '{print $2}'")
	if err!=nil {
		
	}
	fmt.Println(str)
	strarry:=strings.Split(strings.TrimSpace(str),"\n")
	fmt.Println(len(strarry))
	fmt.Println(strarry)
	return
	files, dirs, _ := GetFilesAndDirs("/proc")

	for _, dir := range dirs {
		fmt.Printf("获取的文件夹为[%s]\n", dir)
	}

	for _, table := range dirs {
		temp, _, _ := GetFilesAndDirs(table)
		for _, temp1 := range temp {
			files = append(files, temp1)
		}
	}


	fmt.Printf("=======================================\n")
	xfiles, _ := GetAllFiles("./simplemath")
	for _, file := range xfiles {
		fmt.Printf("获取的文件为[%s]\n", file)
	}
}

//获取指定目录下的所有文件和目录
func GetFilesAndDirs(dirPth string) (files []string, dirs []string, err error) {
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, nil, err
	}

	PthSep := string(os.PathSeparator)
	//suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

	for _, fi := range dir {
		if fi.IsDir() { // 目录, 递归遍历
			//fmt.Println(dirPth+string(os.PathSeparator)+fi.Name())
			dirs = append(dirs, dirPth+PthSep+fi.Name())
			absPath:=dirPth+string(os.PathSeparator)+fi.Name()
		//	fmt.Println(fi.Name())
			IsNodeProcess(absPath+string(os.PathSeparator)+"status",fi.Name())
		///	GetFilesAndDirs(dirPth + PthSep + fi.Name())
		} else {
			// 过滤指定格式
		//	ok := strings.HasSuffix(fi.Name(), ".go")
		//	if ok {
		//		files = append(files, dirPth+PthSep+fi.Name())
		//	}
		}
	}

	return files, dirs, nil
}
func IsNodeProcess(path string,midPath string)(bool){
	bools:=PathExistsWithResult(path)
	if bools {
		raw, err := ioutil.ReadFile(path)

		self := new(ProcessStatus)
		fmt.Sscanf(string(raw), "%s",
			&self.Name)
		fmt.Println(self.Name)
	//	fmt.Println(string(raw))
		if err!=nil {

		}
	}

	return bools
}
//获取指定目录下的所有文件,包含子目录下的文件
func GetAllFiles(dirPth string) (files []string, err error) {
	var dirs []string
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}

	PthSep := string(os.PathSeparator)
	//suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

	for _, fi := range dir {
		if fi.IsDir() { // 目录, 递归遍历
			dirs = append(dirs, dirPth+PthSep+fi.Name())
			GetAllFiles(dirPth + PthSep + fi.Name())
		} else {
			// 过滤指定格式
			//ok := strings.HasSuffix(fi.Name(), ".go")
			//if ok {
			//	files = append(files, dirPth+PthSep+fi.Name())
			//}
		}
	}

	// 读取子目录下文件
	for _, table := range dirs {
		temp, _ := GetAllFiles(table)
		for _, temp1 := range temp {
			files = append(files, temp1)
		}
	}

	return files, nil
}

func maisn() {
	files, dirs, _ := GetFilesAndDirs("/proc")

	for _, dir := range dirs {
		fmt.Printf("获取的文件夹为[%s]\n", dir)
	}

	for _, table := range dirs {
		temp, _, _ := GetFilesAndDirs(table)
		for _, temp1 := range temp {
			files = append(files, temp1)
		}
	}

	for _, table1 := range files {
		fmt.Printf("获取的文件为[%s]\n", table1)
	}

	fmt.Printf("=======================================\n")
	xfiles, _ := GetAllFiles("./simplemath")
	for _, file := range xfiles {
		fmt.Printf("获取的文件为[%s]\n", file)
	}
}


func PathExistsWithResult(path string)(bool){
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}