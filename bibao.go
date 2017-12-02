package main

import (
	"fmt"
)

type System interface {
	play(what string) string
	work(what string) string
}

type Windows struct {

}

func (windows Windows) play(what string) string {
	return what
}

func (windows Windows) work(what string) string {
	return what
}


func main() {
	var system System
	system = new (Windows)
	var a = system.play("wow")
	var b = system.work("code")
	//fmt.Println(a,b)
	fmt.Printf("玩%s,写%s",a,b)

}