package test

import (
	"sync"
	"testing"
)


// 测试一下sync库中的东西

func TestSyncCond(t *testing.T) {
	lock := sync.Mutex{}
	cond = sync.NewCond(lock)
	if cond != nil {

	}

}


// 测试文件锁
// func TestRWMutex(t *testing.T) {
//}


// 测试json的文件读写
// func TestJson(t *testing.T) {
// }


