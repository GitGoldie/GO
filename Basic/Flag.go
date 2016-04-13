package main

import ("fmt"; "flag";"strings")

/* note 
	* you have to issue flag.Parse() to update flag variables
	* you have to access to flag variables indirectly through pointer, *var
	* you can parse command-line arguments through flag.Args()
	* minimum packages to import : flag
*/
func main(){

	//flag n, s 
	var n = flag.Bool("n",false,"omit trailing newline")
	var sep = flag.String("s"," ","separator")  //name, default, description for invalid usage
	
	flag.Parse()
	fmt.Println("Your input: ",flag.Args())
	fmt.Println("Your -s :",strings.Join(flag.Args(), *sep))
	
	//if n has specified
	if !*n {
		fmt.Println()
	}
	
	//my curiousity about new function 
	p := new(int)
	fmt.Println(p, *p)
}