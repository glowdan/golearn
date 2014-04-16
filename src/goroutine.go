package main

import (
	"time"
	"fmt"
	"math/rand"
)

func say(s string) {
	for i:=0;i<5;i++ {
		time.Sleep(time.Second)
		fmt.Println(s)
	}
}
func IsReady(what string, minutes int32) {
	time.Sleep(time.Duration(rand.Int31n(minutes)) * time.Second)
	fmt.Println(what, "is ready")
}

func main() {
	say("yes")
	go say("goroutine")
	say("no")
	go IsReady("tea", 2)
	go IsReady("coffee", 1)
	//IsReady("last", 1)
	//var input string
	//fmt.Scanln(&input)//使用此
	fmt.Println("I'm waiting…")
}
