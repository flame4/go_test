package main

import (
	"encoding/json"
	"os"
	"strconv"
	"sync"
)


func main() {
 	testjsonrw()
	testmultijson()

}

// 测试json的文件读写使用.
func testjsonrw() {
	file1 := "/tmp/test_json1"
	file2 := "/tmp/test_json2"
	file3 := "/tmp/test_json3"

	json_string := "{\"name\": 123, \"position\": \"do not let you go!\"}"
	json_byte := []byte(json_string)

	// copy also work, but if you specific length of a byte slice
	// it will fill in other bit with 0, which is not convenient.
	// copy(json_byte, json_string)

	json_map := make(map[string]interface{})
	json_map["aa"] = 1
	json_map["bb"] = "hello"

	// b1, err := json.Marshal(json_map); 	ErrorClient(err)

	in1, err := os.Create(file1); ErrorClient(err); defer in1.Close()
	in1.Write(json_byte)

	in2, err := os.Create(file2); ErrorClient(err); defer in2.Close()
	in2.WriteString(json_string)

	in3, err := os.Create(file3); ErrorClient(err); defer in3.Close()
	encoder := json.NewEncoder(in3)
	encoder.Encode(json_map)
}

// 测试多个go程一起读写文件的阻塞, 使用文件读写锁.
func testmultijson() {
	mutex := new(sync.Mutex)
	c := make(chan bool, 100)
	for i:=0; i<100; i++ {
		go mj(i, &mutex, c)
	}
	// go的主函数结束后, 不会等待go程序结束, 而是直接退出.
	// go的阻塞似乎没有类似join的玩意儿? 用管道进行阻塞.
	for i:=0; i<100; i++ {
		<- c
	}
}

func mj(i int, mutex *sync.Mutex, c chan bool){
	// defer c <- true: syntax error
	// defer func(c chan bool) {c <- true}: syntax error
	defer
	path := "/tmp/mj"

	j := make(map[string]interface{})
	j["id"+strconv.Itoa(i)] = i
	j["content"+strconv.Itoa(i)] = "json" + strconv.Itoa(i)

	mutex.Lock()
	defer mutex.Unlock()
	if _, err := os.Stat(path); os.IsNotExist(err) {
		file, err := os.Create(path); ErrorClient(err); defer file.Close()
		encoder := json.NewEncoder(file)
		encoder.Encode(j)
		return
	} else {
		file, err := os.Open(path); ErrorClient(err); defer file.Close()
		decoder := json.NewDecoder(file)
		decoder.Decode(j)
		encoder := json.NewEncoder(file)
		encoder.Encode(j)
	}
}
