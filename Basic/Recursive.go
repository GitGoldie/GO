package main

import "fmt"

func main(){
	fmt.Printf("fib last %d\n",fib(20))
}

func fib(x int) int {
	fmt.Printf("fib %d\n",x)
	
	if (x < 2){ return x }
	
	return fib(x-1) + fib(x-2)
}