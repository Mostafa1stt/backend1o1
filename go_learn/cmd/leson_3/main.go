package main

import(
	"fmt"
)


func main(){
	var intarray =[...]int{1,2,3}
	intarray[1] = 123
	fmt.Println(intarray[1:3])
	fmt.Println(&intarray[1])

	//slicing
	var intslice []int =[]int{1,2,3}
	fmt.Printf("the len %v  the capcity %v \n",len(intslice),cap(intslice))
	intslice = append(intslice, 7)
	fmt.Printf("the len %v  the capcity %v \n",len(intslice),cap(intslice))
	var intslice2 []int =[]int{1,2,3}
	intslice = append(intslice, intslice2...)
	fmt.Println(intslice2)

	//make(type,len,cap)
	var intslice3 []int =make([]int,3,8)
	fmt.Println(intslice3)

	//map
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var mymap2 = map[string]int {"adam":23,"sarah":45}
	fmt.Println(mymap2["adam"])
	fmt.Println(mymap2)
	var age,ok = mymap2["neon"]
	if ok==true{
		fmt.Println(age)
	}else{
		fmt.Println("Invalid name")
	}

	for name,age := range mymap2{
		fmt.Printf("Name: %v , Age: %v \n",name,age)
	}

	for i:=0;i<10;i++{
		fmt.Println(i)
	}
	
}
