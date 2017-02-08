package main


import (
	"log"
)



func ErrorClient(err error) bool{
	return Error(err)
}

func Error(err error) bool{
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}