package main

import (
	"bytes"
	"container/list"
	"reflect"
	"time"

	"os/exec"
	"strings"
)

func getNodejsProcessPid() []string {
	getNodePidCollections := "ps -ef | grep  node  | grep -v grep | awk '{print $2}'"
	raw_node, err_node := getProcessPidByName(getNodePidCollections)
	if err_node != nil {

	}

	getPm2PidCollections := "ps -ef | grep  pm2  | grep -v grep | awk '{print $2}'"
	raw_pm2, err_pm2 := getProcessPidByName(getPm2PidCollections)
	if err_pm2 != nil {

	}
	node_arr := strings.Split(strings.TrimSpace(raw_node), "\n")
	pm2_arr := strings.Split(strings.TrimSpace(raw_pm2), "\n")

	result_arr := mergeArray(node_arr, pm2_arr)
	return result_arr
}

func getMapProcess() map[string]string {

	var pidMaps map[string]string
	pidMaps = make(map[string]string)
	pidArray := getNodejsProcessPid()
	for i := 0; i < len(pidArray); i++ {
		pidMaps[pidArray[i]] = pidArray[i]
	}
	return pidMaps
}

func getNodeAgentCount() int {
	getNodePidCollections := "ps -ef | grep  NodeAgent  | grep -v grep | awk '{print $2}'"
	raw_node, err_node := getProcessPidByName(getNodePidCollections)
	node_arr := strings.Split(strings.TrimSpace(raw_node), "\n")
	if err_node != nil {

	}
	return len(node_arr)
}

func mergeArray(strarray1 []string, strarray2 []string) []string {
	length := len(strarray1) + len(strarray2)
	//var result [length]string
	result := make([]string, length)
	for i := 0; i < len(strarray1); i++ {
		result[i] = strarray1[i]
	}
	for i := len(strarray1); i < length; i++ {
		result[i] = strarray2[i-len(strarray1)]
	}
	return result
}

func convertArrayToMap(arraystr []string) map[string]string {
	var convertMap map[string]string
	convertMap = make(map[string]string)
	for i := 0; i < len(arraystr); i++ {
		convertMap[arraystr[i]] = arraystr[i]
	}
	return convertMap

}

func calculateAddPid(globalMap map[string]string, pidMap map[string]string) []string {
	globalLength := len(globalMap)
	pidMaplength := len(pidMap)
	resultList := list.New()
	if globalLength == pidMaplength {
		for key, _ := range pidMap {
			_, ok := globalMap[key]
			if ok {

			} else {
				resultList.PushBack(key)
			}
		}
		return convertListToArray(*resultList)
	}
	return nil
}

func calculateMissPid(globalMap map[string]string, pidMap map[string]string) (string, []string) {
	globalLength := len(globalMap)
	pidMaplength := len(pidMap)
	resultList := list.New()
	if globalLength == pidMaplength {

		for key, _ := range globalMap {
			_, ok := pidMap[key]
			if ok {

			} else {
				resultList.PushBack(key)
			}
		}
		return "UPDATEDOWNPID", convertListToArray(*resultList)
	}

	if globalLength > pidMaplength {
		for key, _ := range globalMap {
			_, ok := pidMap[key]
			if ok {

			} else {
				resultList.PushBack(key)
			}
		}
		return "DOWN", convertListToArray(*resultList)
	}
	if globalLength < pidMaplength {
		for key, _ := range pidMap {
			_, ok := globalMap[key]
			if ok {

			} else {
				resultList.PushBack(key)
			}
		}
		return "ADD", convertListToArray(*resultList)
	}

	return "", nil
}

func convertListToArray(resultList list.List) []string {
	var count int = 0

	resultArray := make([]string, resultList.Len())
	for e := resultList.Front(); e != nil; e = e.Next() {
		resultArray[count] = (string)(e.Value.(string))
		count++
	}
	return resultArray
}

func getProcessPidByName(name string) (string, error) {
	//函数返回一个*Cmd，用于使用给出的参数执行name指定的程序
	cmd := exec.Command("/bin/bash", "-c", name)

	//读取io.Writer类型的cmd.Stdout，再通过bytes.Buffer(缓冲byte类型的缓冲器)将byte类型转化为string类型(out.String():这是bytes类型提供的接口)
	var out bytes.Buffer
	cmd.Stdout = &out

	//Run执行c包含的命令，并阻塞直到完成。  这里stdout被取出，cmd.Wait()无法正确获取stdin,stdout,stderr，则阻塞在那了
	err := cmd.Run()
	//checkErr(err)

	return out.String(), err
}

func exeCommand(command string) (string, error) {
	//函数返回一个*Cmd，用于使用给出的参数执行name指定的程序
	cmd := exec.Command("/bin/bash", "-c", command)

	//读取io.Writer类型的cmd.Stdout，再通过bytes.Buffer(缓冲byte类型的缓冲器)将byte类型转化为string类型(out.String():这是bytes类型提供的接口)
	var out bytes.Buffer
	cmd.Stdout = &out

	//Run执行c包含的命令，并阻塞直到完成。  这里stdout被取出，cmd.Wait()无法正确获取stdin,stdout,stderr，则阻塞在那了
	err := cmd.Run()
	//checkErr(err)

	return out.String(), err
}
func getProcessStartTime(pid string) string {
	str := "ps -p " + pid + "  -o lstart"
	result, err := exeCommand(str)
	if err != nil {
		return ""
	}
	node_arr := strings.Split(strings.TrimSpace(result), "\n")

	if len(node_arr) > 1 {
		t, err := time.Parse(time.ANSIC, node_arr[1])
		if err != nil {
			return ""
		}
		return t.String()
	}
	return ""
}

func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}
