package main

import (
	"encoding/json"
	"os"

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
	json_byte = []byte(json_string)

	// copy also work, but if you specific length of a byte slice
	// it will fill in other bit with 0, which is not convenient.
	// copy(json_byte, json_string)
	json_map := make(map[string][interface{}])
	json_map["aa"] = 1
	json_map["bb"] = "hello"

	b1, err := json.Marshal(json_map); 	ErrorClient(err)

	in, err := os.Create(file1); ErrorClient(err)

}

func testmultijson() {

}
