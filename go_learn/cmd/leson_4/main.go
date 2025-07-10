package main

import(
"fmt"
"strings"
)
func main(){
	var mystring = "résumé"
	var indeed = mystring[0]
	fmt.Printf("%v , %T\n",indeed,indeed)
	for i, v := range mystring{
		fmt.Println(i,v)
	}

	var  mystring2= []rune("résumé")
	var indeed2 = mystring[0]
	fmt.Printf("%v , %T\n",indeed2,indeed2)
	for i, v := range mystring2{ 
		fmt.Println(i,v)
	}
	fmt.Printf("the len of the the string is %v \n ", len(mystring2))

	var veigo = 'é'
	fmt.Printf("the type  of the the string is %T and the value is %v \n", veigo,veigo)

	var mystring3 = []string{"m","o","s","t","a","f","a"}
	var mystring4 = ""
	for i := range mystring3{
	mystring4 += mystring3[i]
	}
	fmt.Println(mystring4)

	var mystring5 = []string{"m","o","s","t","a","f","a"}
	var mystring6 strings.Builder
	for i := range mystring5{
	mystring6.WriteString(mystring5[i])
	}
	var mystring7 = mystring6.String()
	fmt.Printf("value : %v type : %T",mystring7,mystring7)

}
