/*
 * This file is for : print format 
 */
package main

import ("fmt")

func main(){

	// const and auto type assignment
	const max = 10		
	
	// display type 
	fmt.Println("\n=======================================")
	fmt.Println("Same Variable Can be displayed as different format\n%d\t%b\t%x\t%q\t%T\t")
	fmt.Println("=======================================")
	for i := 0; i < max; i++ {
		fmt.Printf("%d\t%b\t%x\t%q\t%T\n",i,i,i,i,i)
	}
	fmt.Println("=======================================")
}

/*
 * How to generate execution file : go build <filename>
 */
