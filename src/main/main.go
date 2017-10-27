package main

import (
	"fmt"
	"sync"
)

type aa struct {
	a int
	b int
}

func main() {
 	//testjsonrw()
	//testmultijson()
	//testjsondecode()

	x := sync.WaitGroup{}
	x.Add(5)
	t := 0
	for i:=0; i<5; i++ {
		go func(t int){
			t++
			fmt.Println(t)
			x.Done()
		}(t)
	}
	x.Wait()





}

