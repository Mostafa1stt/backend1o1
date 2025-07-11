package main 

import (
	"fmt"
)

// type GasEngine struct{
// 	clinder int
// 	horsePower int
// 	 owner
// 	 int
// } 
//
// type owner struct{
// 	name string
// }
//
// func main(){
//
// 	var myengine GasEngine = GasEngine{clinder: 4,horsePower: 600, owner:owner {"mostafa"}, int: 3   } //==  GasEngine{ 4, 600, owner{mostafa} , 3 } 
// 	myengine.clinder = 6
// 	fmt.Println(myengine.clinder,myengine.horsePower,myengine.name,myengine.int)
//
// }

type GasEngine struct{
	clinder int
	horsePower int
} 

type elecEngine struct{
	motors int
	horsePower int
} 

func (e GasEngine) horsepercylinder() int{
	return e.horsePower/e.clinder
}

func (e elecEngine) horsepercylinder() int{
	return e.horsePower/e.motors
}

type engine interface{
	horsepercylinder() int
}

func strongornot(e engine,powerneeded int){
	if powerneeded <= e.horsepercylinder(){
		fmt.Println("strong as fuckkkkkkkkkk")
	}else{
		fmt.Println("weak as fuckkkkkkkkkkk")
	}
}

func main(){

	var myengine1 elecEngine= elecEngine{motors: 4,horsePower: 600} //==  GasEngine{ 4, 600} 
	var myengine2 GasEngine= GasEngine{clinder: 4,horsePower: 600} //==  GasEngine{ 4, 600} 

	fmt.Println(myengine1.motors,myengine1.horsePower)
	fmt.Println(myengine1.horsepercylinder())
	strongornot(myengine1,100)

	fmt.Println(myengine2.clinder,myengine2.horsePower)
	fmt.Println(myengine2.horsepercylinder())
	strongornot(myengine2,100)
}
