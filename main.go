package main

import (
"fmt"
)

//Normal Function
func add(x int, y int) int { //add(x, y int) 
    return x + y
}

//function with multiple return values
func multipleReturnValues( x int) (int,int){
	return x+1, x-1
}

//funtion with named return values
func namedReturnType(x int) (add, mul int){
	add=x+1
	mul = x*x
	return // no need to specify the 
}

//Vardiac function
func sum( Name string, nums ...int) {
	fmt.Println("name = ",Name)
    fmt.Print(nums, " ")
    total := 0
    for _, num := range nums {
        total += num
    }
    fmt.Println(total)
}

func makeEvenGenerator(i int) func(i int) int {
 
  return func(i int) (ret int) {
    ret = i
    i += 2
    return
  }
}


func main(){
	fmt.Println("hello World")
	fmt.Println(add(42, 13))
	add, sub := multipleReturnValues(5)
	fmt.Println("add = ",add," sub = ",sub)
	fmt.Println(namedReturnType(5))
	sum("AAA",1,2,3,4,5)
	
	nums := []int{1, 2, 3, 4}
    sum("BBB",nums...)
    
    nextEven := makeEvenGenerator(2)
  fmt.Println(nextEven(2)) // 0
  fmt.Println(nextEven(2)) // 2
  fmt.Println(nextEven(2)) // 4
	
	
}



