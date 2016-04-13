/*
 * Universal Greeting 
 */
 
package main

//note : unused package will throw an error 
//there are few tricks to cheat and you can search internet
import ("fmt")

func main(){
	var name = "world"
	message = "Hello"
	
	//no need to specify format and include new line
	fmt.Println(name,message)
	
	// if you want to specify format
	fmt.Printf("%s %s\n",name,message)
}
