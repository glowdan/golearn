package main

import (
	"fmt"
	"strings"
)

func main() {
    phrase := "我是a 0的√ sf"
    fmt.Println("index rune char bytes")
    for index,char := range phrase {//索引 码点 字符 字节
        fmt.Printf("%-2d	%U	'%c'	%X\n", index, char, char, []byte(string(char)))
    }
	i := strings.Index(phrase, " ")//使用strings的函数来操作
	firstWord := phrase[:i]
	j := strings.LastIndex(phrase, " ")
	lastWord := phrase[j:]
	fmt.Println(firstWord, lastWord)
	madian := []rune(phrase)//转换成码点
	for _,ma := range madian {
		fmt.Printf("%c\n", ma)
	}
	printString()

}

func IntForBool(b bool) int {
	if b {
		return 1
	} else {
		return 0
	}
}

func printString() {

	fmt.Printf("%t %t \n", true, false)//格式化布尔值
	fmt.Printf("%d %d \n", IntForBool(true), IntForBool(false))//格式化布尔值
	//数字
	fmt.Printf("%b|%9b|%-9b|%09b|% 9b\n", 37,37,37,37,37)
	fmt.Printf("%o|%#b|%#8od|%#+ 8o|%+08o\n", 41,41,41,41,-41)
	//string 都是rune
	fmt.Printf("%d %#04X %U '%c'", '我', 935, '\u0564', '\U000003A6')
	//切片的格式
	runeString := "q 切片de √ / "
	fmt.Printf("%s\n%q\n%+q\n%#q\n%x\n%X\n%#X\n", runeString, runeString, runeString, runeString, []rune(runeString), []rune(runeString), []rune(runeString))
	stringsPkg()
}


func stringsPkg() {
	fmt.Println(strings.Contains("我的a", "a"))
}
