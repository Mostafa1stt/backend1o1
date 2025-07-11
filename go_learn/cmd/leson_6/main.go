package main

import(
	"fmt"
)

func main(){
	// var p *int = new(int)
	// i := 2 
	// fmt.Printf("the of the thing that p points to is : %v \n",*p)
	// fmt.Printf("the value of it self ie(the address it is pointing to) : %v \n",p)
	// fmt.Printf("the value of i is : %v \n",i)
	// fmt.Printf("the address of i is : %v \n",&i)
	// p = &i
	// fmt.Printf("after p = &i \n")
	// fmt.Printf("the of the thing that p points to is : %v \n",*p)
	// fmt.Printf("the value of it self ie(the address it is pointing to) : %v \n",p)
	// fmt.Printf("the value of i is : %v \n",i)
	// fmt.Printf("the address of i is : %v \n",&i)

	// var slice = []int{1,2,3}
	// var slicecopy = slice
	// slice[2] = 8
	// fmt.Println(slice)
	// fmt.Println(slicecopy)

	var thing1 = [5]float64{1,2,3,4,5}
	fmt.Printf("the memory loacation of the thing1 array is %p\n",&thing1)
	var sequared [5]float64 = square(&thing1)
	fmt.Printf("the reseult is : %v\n",sequared)
	fmt.Printf("the reseult is : %v\n",thing1)
}
func square(thing2 *[5]float64) [5]float64{
	fmt.Printf("the memory loacation of the thing2 array is %p\n",thing2)
	for i := range thing2{
		thing2[i] = thing2[i]*thing2[i]

	}
	return *thing2
}

