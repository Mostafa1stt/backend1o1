package main

import (
	"errors"
	"fmt"
)

func main() {
	var printvalue string = "hello world"
	printme(printvalue)

	number1, number2 := 50, 0
	var  div, remin ,error = div2num(number1, number2)
	if error!=nil{
		println(error.Error())
	}else{
		fmt.Printf("bla bla bla %v bla bla bla %v bla bla bla", div, remin)
	}
}

func printme(thestring string) {
	fmt.Println(thestring)
}

func div2num(num1 int, num2 int) (int, int,error) {
	var err error 
	if num2==0{
		err = errors.New("Cannot Divide By Zero")
		return 0,0,err
	}
	return (num1 / num2), (num1 % num2),err
}	

