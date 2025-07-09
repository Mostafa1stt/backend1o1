package main

import "fmt"
import "unicode/utf8"

func main() {
	var num int = 54
	fmt.Println("num:")
	fmt.Println(num)

	var floatnum float32 = 10.1
	var resualt float32 = floatnum + float32(num)
	fmt.Println("res_num:")
	fmt.Println(resualt)

	var num2 int = 72
	fmt.Println("num/:", num2)
	fmt.Println("num/:")
	fmt.Println(num / num2)
	fmt.Println("num%:")
	fmt.Println(num % num2)

	var string string = "hello" + " " + "world"
	fmt.Println(string)

	fmt.Println(utf8.RuneCountInString("γ"))

	var viego rune = 'a'
	fmt.Println(viego)

	var boolen bool = false
	fmt.Println(boolen)

	fansy, fansy2 := 1, 2
	fmt.Println(fansy, fansy2)

}
